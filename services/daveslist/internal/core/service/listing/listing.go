package listing

import "daveslist/internal/core/port"

type Config struct {
	AuthService      port.AuthService
	ListingRepo      port.ListingRepository
	ReplyListingRepo port.ReplyListingRepository
}

type Service struct {
	authService      port.AuthService
	listingRepo      port.ListingRepository
	replyListingRepo port.ReplyListingRepository
}

func New(cfg *Config) port.ListingService {
	return &Service{
		authService:      cfg.AuthService,
		listingRepo:      cfg.ListingRepo,
		replyListingRepo: cfg.ReplyListingRepo,
	}
}
