package services

func (s *service) DeleteMail(mailid uint, emailID string) error {
	err := s.model.DeleteMail(mailid, emailID)
	if err != nil {
		return err
	}
	return nil

}
