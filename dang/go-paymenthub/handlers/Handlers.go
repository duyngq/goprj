package handlers

import (
	"github.com/gin-contrib/opengintracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
	"paymenthub/services"
)

var paymentService = new(services.PaymentService)
var accountService = new(services.AccountService)

func Routers() *gin.Engine {
	//Config jaeger communicate with api gateway
	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	trace, closer := jaeger.NewTracer(
		"api_gateway",
		jaeger.NewConstSampler(true),
		jaeger.NewNullReporter(),
		jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	defer closer.Close()
	opentracing.SetGlobalTracer(trace)
	var fn opengintracing.ParentSpanReferenceFunc
	fn = func(sc opentracing.SpanContext) opentracing.StartSpanOption {
		return opentracing.ChildOf(sc)
	}

	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	client := r.Group("/api")
	client.POST("/account", accountService.StoreAccount)
	client.POST("/inquiry",opengintracing.SpanFromHeadersHttpFmt(trace, "inquiry", fn, false), paymentService.GetInquiry)
	client.POST("/charge", opengintracing.SpanFromHeadersHttpFmt(trace, "charge", fn, false),paymentService.ChargePayment)
	client.POST("/settlement", opengintracing.SpanFromHeadersHttpFmt(trace, "settlement", fn, false),paymentService.SetSettlement)
	client.POST("/void",opengintracing.SpanFromHeadersHttpFmt(trace, "void", fn, false), paymentService.CancelPayment)
	return r
}
