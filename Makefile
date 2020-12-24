mockgen:
	mockgen -package=mocks -source=dto/go -destination=dto/mocks/mock_go
	mockgen -package=mocks -source=usecase/usecase.go -destination=usecase/mocks/mock_usecase.go