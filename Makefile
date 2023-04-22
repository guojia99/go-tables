all:
	go run main.go


tools:
	echo "create colors"
	go run tools/quict_create_colors.go
	echo "create mock struct"
	go run tools/quict_create_mock_struct.go