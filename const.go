package main

const (
	ruleBag = `
		rule CheckPackageTypeV1 "Check package type v1" {
			when
				Package.Module == "Bag"
			then
				CheckBag();
		}

		rule CheckBag "Check Bag" {
			when
				Package.Type == 1
			then
				CheckValidation();
		}

		rule CheckValidation "Check Validation" {
			when
				Package.RunValidationWarehouseCheck() == true && Package.RunValidationShipper() == true
			then
				Package.Complete();
				Retract("CheckPackageTypeV1");
		}
	`
)