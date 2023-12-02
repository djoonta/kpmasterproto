PROTOC_MAIN = ./master.proto
PROTOC_OCCUPATION = ./occupation/occupation.proto
PROTOC_INSTALLMENT = ./installment/installment.proto
PROTOC_ENCUMBRANCE = ./encumbrance/encumbrance.proto
PROTOC_BANK = ./bank/bank.proto
PROTOC_PAGINATION = ./pagination/pagination.proto
PROTOC_VERIFICATION_STATUS_REJECT = ./verification_status_reject/verification_status_reject.proto
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


encumbrance:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_ENCUMBRANCE)

bank:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_BANK)

verification_status_reject:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_VERIFICATION_STATUS_REJECT)

pagination:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_PAGINATION)



all: verification_status_reject bank encumbrance installment occupation pagination main

.PHONY: verification_status_reject bank encumbrance installment occupation pagination main all
