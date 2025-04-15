package efi

import (
	"fmt"
)


type TransactionType string

const (
	
	TransactionTypeCharge    TransactionType = "CHARGE"     
	TransactionTypeDueCharge TransactionType = "DUE_CHARGE" 
	TransactionTypePixSend   TransactionType = "PIX_SEND"   
	TransactionTypeRefund    TransactionType = "REFUND"     

	
	StatusChargeActive        = "ATIVA"
	StatusChargeCompleted     = "CONCLUIDA"
	StatusChargeRemovedByUser = "REMOVIDA_PELO_USUARIO_RECEBEDOR"
	StatusChargeRemovedByPSP  = "REMOVIDA_PELO_PSP"

	
	StatusPixSendProcessing = "EM_PROCESSAMENTO"
	StatusPixSendCompleted  = "REALIZADO"
	StatusPixSendFailed     = "NAO_REALIZADO"

	
	StatusRefundProcessing = "EM_PROCESSAMENTO"
	StatusRefundCompleted  = "DEVOLVIDO"
	StatusRefundFailed     = "NAO_REALIZADO"
)


type TransactionStatus struct {
	ID          string          
	Type        TransactionType 
	Status      string          
	IsCompleted bool            
	IsFailed    bool            
	Message     string          
}


func (c *Client) verifyChargeStatus(status *TransactionStatus) error {
	charge, err := c.ImmediateCharge().GetCharge(status.ID, 0)
	if err != nil {
		return fmt.Errorf("failed to get charge status: %w", err)
	}

	status.Status = charge.Status
	status.IsCompleted = charge.Status == StatusChargeCompleted
	status.IsFailed = charge.Status == StatusChargeRemovedByUser || charge.Status == StatusChargeRemovedByPSP

	return nil
}


func (c *Client) verifyDueChargeStatus(status *TransactionStatus) error {
	charge, err := c.DueCharge().Get(status.ID, 0)
	if err != nil {
		return fmt.Errorf("failed to get due charge status: %w", err)
	}

	status.Status = charge.Status
	status.IsCompleted = charge.Status == StatusChargeCompleted
	status.IsFailed = charge.Status == StatusChargeRemovedByUser || charge.Status == StatusChargeRemovedByPSP

	return nil
}


func (c *Client) verifyPixSendStatus(status *TransactionStatus) error {
	
	pixSend, err := c.PixSend().GetByIDEnvio(status.ID)
	if err != nil {
		
		pixSend, err = c.PixSend().GetByE2EID(status.ID)
		if err != nil {
			return fmt.Errorf("failed to get Pix send status: %w", err)
		}
	}

	status.Status = pixSend.Status
	status.IsCompleted = pixSend.Status == StatusPixSendCompleted
	status.IsFailed = pixSend.Status == StatusPixSendFailed

	return nil
}



func (c *Client) verifyRefundStatus(status *TransactionStatus) error {
	
	e2eID, refundID, ok := parseRefundID(status.ID)
	if !ok {
		return fmt.Errorf("invalid refund ID format, expected 'e2eID:refundID'")
	}

	refund, err := c.PixManagement().GetRefund(e2eID, refundID)
	if err != nil {
		return fmt.Errorf("failed to get refund status: %w", err)
	}

	status.Status = refund.Status
	status.IsCompleted = refund.Status == StatusRefundCompleted
	status.IsFailed = refund.Status == StatusRefundFailed

	return nil
}


func parseRefundID(combinedID string) (e2eID, refundID string, ok bool) {
	for i, c := range combinedID {
		if c == ':' {
			e2eID = combinedID[:i]
			refundID = combinedID[i+1:]
			return e2eID, refundID, e2eID != "" && refundID != ""
		}
	}
	return "", "", false
}


func (s *TransactionStatus) setMessage() {
	switch s.Type {
	case TransactionTypeCharge, TransactionTypeDueCharge:
		switch s.Status {
		case StatusChargeActive:
			s.Message = "Charge is active and ready for payment"
		case StatusChargeCompleted:
			s.Message = "Charge has been paid successfully"
		case StatusChargeRemovedByUser:
			s.Message = "Charge was removed by the receiving user"
		case StatusChargeRemovedByPSP:
			s.Message = "Charge was removed by the payment service provider"
		default:
			s.Message = fmt.Sprintf("Unknown charge status: %s", s.Status)
		}
	case TransactionTypePixSend:
		switch s.Status {
		case StatusPixSendProcessing:
			s.Message = "Pix send is being processed"
		case StatusPixSendCompleted:
			s.Message = "Pix was sent successfully"
		case StatusPixSendFailed:
			s.Message = "Pix send failed"
		default:
			s.Message = fmt.Sprintf("Unknown Pix send status: %s", s.Status)
		}
	case TransactionTypeRefund:
		switch s.Status {
		case StatusRefundProcessing:
			s.Message = "Refund is being processed"
		case StatusRefundCompleted:
			s.Message = "Refund was completed successfully"
		case StatusRefundFailed:
			s.Message = "Refund failed"
		default:
			s.Message = fmt.Sprintf("Unknown refund status: %s", s.Status)
		}
	}
}
