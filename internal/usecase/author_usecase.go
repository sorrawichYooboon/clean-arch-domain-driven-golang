package usecase

import (
	"fmt"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/apperror"
)

type AuthorUseCaseImpl struct {
	authorRepo repository.AuthorRepository
	cacheRepo  repository.CacheAuthorRepository
}

func NewAuthorUseCase(authorRepo repository.AuthorRepository, cacheRepo repository.CacheAuthorRepository) AuthorUseCase {
	return &AuthorUseCaseImpl{
		authorRepo: authorRepo,
		cacheRepo:  cacheRepo,
	}
}

func (uc *AuthorUseCaseImpl) GetAllAuthors() ([]domain.Author, error) {
	resp := []domain.Author{}
	authors, err := uc.cacheRepo.GetAll()
	if err != nil {
		return nil, &apperror.ErrCacheDatabase
	}

	if len(authors) > 0 {
		for _, author := range authors {
			resp = append(resp, mapToDomainAuthor(author))
		}
		return resp, nil
	}

	authors, err = uc.authorRepo.GetAll()
	if err != nil {
		return nil, &apperror.ErrDatabase
	}

	if err := uc.cacheRepo.SetAll(authors); err != nil {
		return nil, &apperror.ErrCacheDatabase
	}

	for _, author := range authors {
		resp = append(resp, mapToDomainAuthor(author))
	}

	return resp, nil
}

func (uc *AuthorUseCaseImpl) GetAuthorByID(id uint) (*domain.Author, error) {
	author, err := uc.authorRepo.GetByID(id)
	if err != nil {
		return nil, &apperror.ErrDatabase
	}

	resp := mapToDomainAuthor(*author)

	return &resp, nil
}

func (uc *AuthorUseCaseImpl) CreateAuthor(name, bio string) error {
	author := domain.NewAuthor(name, bio)

	createAuthorReq := &repository.Author{
		Name: author.Name,
		Bio:  author.Bio,
	}

	err := uc.authorRepo.Create(createAuthorReq)
	if err != nil {
		return &apperror.ErrDatabase
	}

	authors, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		cacheAuthor := mapToRepoAuthor(*author)

		authors = append(authors, cacheAuthor)
		if cacheErr := uc.cacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *AuthorUseCaseImpl) UpdateAuthor(author *domain.Author) error {
	updateAuthorReq := &repository.Author{
		ID:   author.ID,
		Name: author.Name,
		Bio:  author.Bio,
	}

	err := uc.authorRepo.Update(updateAuthorReq)
	if err != nil {
		return &apperror.ErrDatabase
	}

	authors, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, a := range authors {
			if a.ID == author.ID {
				cacheAuthor := mapToRepoAuthor(*author)

				authors[i] = cacheAuthor
				break
			}
		}
		if cacheErr := uc.cacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *AuthorUseCaseImpl) DeleteAuthor(id uint) error {
	err := uc.authorRepo.Delete(id)
	if err != nil {
		return &apperror.ErrDatabase
	}

	authors, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, a := range authors {
			if a.ID == id {
				authors = append(authors[:i], authors[i+1:]...)
				break
			}
		}
		if cacheErr := uc.cacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func mapToDomainAuthor(repoAuthor repository.Author) domain.Author {
	return domain.Author{
		ID:   repoAuthor.ID,
		Name: repoAuthor.Name,
		Bio:  repoAuthor.Bio,
	}
}

func mapToRepoAuthor(domainAuthor domain.Author) repository.Author {
	return repository.Author{
		ID:   domainAuthor.ID,
		Name: domainAuthor.Name,
		Bio:  domainAuthor.Bio,
	}
}
