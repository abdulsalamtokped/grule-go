package main

const (
	ruleBag = `
		rule CheckModule "Check package type v1" salience 1 {
			when
				Package.Module == Module
			then
				CheckType();
		}
		rule CheckType "Check Package" salience 2 {
			when
				Package.Type == Type
			then
				CheckValidation();
		}
		rule CheckValidation "Check Validation" salience 3 {
			when
				Package.RunValidationWarehouseCheck() == true && Package.RunValidationShipper() == true
			then
				Package.Complete();
		}
	`
)
