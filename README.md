# NumX Library

NumX is a simple Go library for generating unique IDs and One-Time Passwords (OTPs). This document provides instructions on how to install and use the library.

## How To Install

To install the NumX library, run the following command:

go get github.com/tss182/numx


## How To Use

Here’s a quick example of how to use the NumX library in your Go application:

// Generate ID From Time
genID := numx.GenerateID()
fmt.Println("Generated OTP:", otp) // will generate ID base on time and random char

// Transform ID to time.Time
timeID := numx.GetTimeFromID(genID) // genID is the ID from numx.GenerateID
fmt.Println("Time from ID:", timeID)

// Generate OTP
otp := numx.GetOtp(4) // 4 is the sample length for OTP
fmt.Println("Generated OTP:", otp) // will generate a 4-digit random number


## Functions Overview

- **GenerateID()**: Generates a unique ID.
- **GetTimeFromID(id string)**: Converts the generated ID back to a `time.Time` object.
- **GetOtp(length int)**: Generates a One-Time Password of the specified length.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
