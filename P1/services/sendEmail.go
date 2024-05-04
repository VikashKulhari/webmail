package services

import (
	"errors"

	"github.com/VikashKulhari/P1/entities"
)

func (s *service) SendEmail(emailReq entities.EmailReq) error {
	var email entities.Email
	if len(emailReq.To) > 0 {

		isToExists := s.model.IsExistingUserByEMailID(emailReq.To)
		if isToExists {
			email.To = emailReq.To
			email.From = emailReq.From
			email.Body = emailReq.Body
			email.Subject = emailReq.Subject
			email.IsDraft = false
			email.IsDeleted = false
			email.IsImportant = false
			email.IsSpam = false
			email.RepliesCount = 0
			// email.Thread = []entities.Replies{}
			err := s.model.SendEmail(email)
			if err != nil {
				return err
			}
			return nil
			// send mail
			//if any error => say internal server error and save in draft and ask for "Retry , saved in drafts"
		} else {
			return errors.New("user you want to send mail does not exist")
			// save in draft with status above reason
		}

	} else {
		email.To = "draft"
		email.From = emailReq.From
		email.Body = emailReq.Body
		email.Subject = emailReq.Subject
		email.IsDraft = true
		email.IsDeleted = false
		email.IsImportant = false
		email.IsSpam = false
		email.RepliesCount = 0
		// email.Thread = []entities.Replies{}
		err := s.model.SendEmail(email)
		if err != nil {
			return errors.New("add reciepient. saved in drafts")
		}
		return errors.New("couldn't save the mail. please add reciepient")
	}
}
