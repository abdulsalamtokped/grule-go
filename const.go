package main

const (
	rulePlain = `
		rule Custom "Custom" salience 1{
			when
				Package.Module == Module
			then
				Log("Ok");
				Package.RunAfter();

				Complete();
		}
	`
	ruleBag = `
		rule CheckModule "Check package type v1" {
			when
				Package.Module == Module
			then
				CheckType();
		}
		rule CheckType "Check Package" {
			when
				Package.Type == Type
			then
				CheckValidation();
		}
		rule CheckValidation "Check Validation" {
			when
				Package.RunValidationWarehouseCheck() == true && Package.RunValidationShipper() == true
			then
				Package.Complete();
		}
	`
)
