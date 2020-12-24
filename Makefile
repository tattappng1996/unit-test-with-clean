mockgen:
	mockgen -package=mocks -source=usecase/usecase.go -destination=usecase/mocks/mock_usecase.go