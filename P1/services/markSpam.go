package services

func (s *service) MarkSpam(mailid uint, emailID string) error {
	err := s.model.MarkSpam(mailid, emailID)
	if err != nil {
		return err
	}
	return nil

}
