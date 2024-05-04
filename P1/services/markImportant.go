package services

func (s *service) MarkImportant(mailid uint, emailID string) error {
	err := s.model.MarkImportant(mailid, emailID)
	if err != nil {
		return err
	}
	return nil

}
