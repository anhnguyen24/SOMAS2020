package server

// probeDisaster checks if a disaster occurs this turn
func (s *SOMASServer) probeDisaster() (bool, error) {
	s.logf("start probeDisaster")
	defer s.logf("finish probeDisaster")
	// TODO:- env team
	return false, nil
}
