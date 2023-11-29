PROTOC_MAIN = ./master.proto
PROTOC_OCCUPATION = ./occupation/occupation.proto
PROTOC_INSTALLMENT = ./installment/installment.proto
PROTOC = protoc
PROTOC_FLAGS = -I . --go_out=. --go_opt=paths=source_relative \
							 --go-grpc_out=. --go-grpc_opt=paths=source_relative

occupation:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_OCCUPATION)

main:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_MAIN)

installment:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_INSTALLMENT)

all: installment occupation main

.PHONY: installment occupation main all
