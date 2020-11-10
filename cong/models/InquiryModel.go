package models

import  ."time"

type Inquiry struct {
	CisId string				`json:"cis_id" `
	CreatedAt  Time			`json:"created_at" `
	UpdatedAt Time 				`json:"updated_at" `
	PointTotal int 				`json:"point_total"  `
	AvailablePoints [] InquiryPoint   	`json:"available_points" `
}

type InquiryPoint struct {
	ExpiredAt Time 				`json:"expired_at"  `
	Point int 					`json:"point"  `
}
