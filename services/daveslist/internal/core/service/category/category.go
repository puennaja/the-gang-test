package category

import "daveslist/internal/core/port"

type Config struct {
	CategoryRepo   port.CategoryRepository
	ListingService port.ListingService
}

type Service struct {
	categoryRepo   port.CategoryRepository
	listingService port.ListingService
}

func New(cfg *Config) port.CategoryService {
	return &Service{
		categoryRepo:   cfg.CategoryRepo,
		listingService: cfg.ListingService,
	}
}
