package main

import (
	"bytes"
	"io/ioutil"
	"kbyp-lambda-go/utilities"
	"net/http"
	"strings"
	"github.com/krazybee/gofpdf"
)

func main() {

	pdfBytes := genPDF()
	ioutil.WriteFile("test2.pdf", pdfBytes, 0777)
}

type bufWriter struct {
	buf []byte
}

func (nw *bufWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	// utilities.Println("pdf write bytes: ", n)
	nw.buf = p
	return
}

func addFooterToPDF(pdf *gofpdf.Fpdf, extaData map[string]string) {
	footerPlacement := "R"
	pdf.SetFooterFuncLpi(func(lastPage bool) {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 6)
		pdf.CellFormat(0, 0, "Page of ", "", 0, footerPlacement, false, 0, "")
		pdf.Ln(0)		
	})
}

func (nw *bufWriter) Close() (err error) {
	return
}

func genPDF() []byte {
	pw := new(bufWriter)
	pdf := gofpdf.New("P", "mm", "A4", "")
	fontsize := 12.0
	smallfontsize := 10.0
	halfWidth := 90.0
	font := "Times New Roman"

	pdf.AddPage()
	addFooterToPDF(pdf, nil)
	pdf.SetMargins(10, 10, 10)
	pdf.SetAutoPageBreak(true, 15)

	pdf.SetFont(font, "I", smallfontsize)
	pdf.CellFormat(0, 10, "Krazybee Legal     ", "", 0, "R", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "BU", fontsize)
	pdf.CellFormat(0, 10, "LOAN APPLICATION FORM", "0", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "ACCOUNT TYPE:", "1", 0, "C", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Personal", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "NAME:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "FATHER/SPOUSE NAME:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "MOTHER'S NAME:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "DATE OF BIRTH:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "GENDER:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "MARITAL STATUS:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "OCCUPATION:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "NATIONALITY:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "RESIDENTIAL STATUS:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "PROOF OF IDENTITY:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "PAN:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "PROOF OF ADDRESS:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "ADDRESS TYPE:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 40, "ADDRESS:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")

	pdf.Ln(0)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 40, "CURRENT ADDRESS:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 6, "L", false, 0, "")
		
	pdf.Ln(0)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "PHONE NUMBER:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "EMAIL:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "NAME OF BANK:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "BANK A/C NO:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "IFSC:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "NAME OF RELATED PERSON:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "PHONE NUMBER OF RELATED PERSON:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")
	pdf.Ln(-1)

	pdf.AddPage()

	pdf.SetFont(font, "B", fontsize)
	pdf.CellFormat(0, 10, "SUMMARY OF THE LOAN TERMS", "0", 0, "C", false, 0, "")
	
	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(20, 10, "S. No.", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "Particulars", "1", 0, "L", false, 0, "")
	pdf.CellFormat(110, 10, "Details", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(20, 10, "1", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "LOAN ID/ SERIAL ID", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(110, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(20, 10, "2", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "CITY", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(110, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(20, 10, "3", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "PROSPECT No.", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(110, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)

	pdf.MultiCell(0, 8, "I understand the terms of the loan to be provide to me, if approved by the Lender as per its internal policies and law shall be as specified below (\"Loan\"):", "0", "L", false)

	//pdf.Ln(0)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "PARTICULARS (Below charges are exclusive of ", "TLR", 0, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "DETAILS", "TLR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 10, "GST).)", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "LR", false, 0, "")

	pdf.Ln(0)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Lender:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Vivriti Capital Private Limited", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Platform:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Loan Amount:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Rate of Interest:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Purpose of Loan:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Platform Credit Reassessment Fees:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Platform Processing Fees:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Platform Service Charges:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Prepayment Charges:", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 10, "", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 8, "(i) In case of prepayment of any", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "instalment/ part prepayment:", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "The amount equivalent to the differential between the", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "(ii) In case of prepayment in full:", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "New Outstanding Balance and the Existing", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "Outstanding Balance shall be paid as prepayment", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "charges by the Borrower in case of prepayment of any", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "instalment/ part prepayment of the Loan.", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "[pre-determined rate to be included].", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, `"Existing Outstanding Balance" means the`, "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "Outstanding Balance which was remaining due to be", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "paid under the Loan prior to prepayment of the", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "relevant instalment under the Loan, as reflected under", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "the repayment schedule provided by the Lenders under", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "Clause 5.5.", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, `"New Outstanding Balance" means the Outstanding`, "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "Balance remaining due under the Loan pursuant to and", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "post prepayment of any instalment under the ", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "Loan.", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, `"Outstanding Balance" means collectively the`, "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "principal, interest, compound interest, default charges/", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "interest, any other charges, dues and monies payable,", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "costs and expenses, reimbursable, as outstanding from", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "time to time and whether any of them are due or not in", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "relation to the Loan.", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 0, "", "1", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 0, "", "1", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Late Payment Charges;", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Default Charges/rate:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "NACH Dishonour Charges:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "Banking Details for disbursal of Loan:", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(halfWidth, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.MultiCell(0, 10, "I agree to submit the following documents for availing of Loan:", "0", "L", false)

	//pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(30, 10, "S. No.", "1", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "Documents from the Applicant and the Co-Applicant (if any)", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 10, "Status", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(30, 10, "1", "1", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "PAN Card or Form 60*", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(30, 10, "2", "1", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "Passport/Voter ID/Masked Aadhaar/Driver License", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(30, 10, "3", "1", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "Last 3 months bank statements or other income proof", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(30, 10, "4", "1", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "Any other document requested by Lender", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, "I acknowledge, understands and agrees that Lender has adopted risk-based pricing, which is arrived by taking into account, broad parameters like the customers financial and credit profile. I also understand all the terms listed above along with certain fees/ charges (as may be applicable) apply for the said Loan from the Lender. I further understand and acknowledge that as may be applicable the Platform Processing Fees, Platform Service Charges, One Time Onboarding Fees, Platform Credit Reassessment Fees, Prepayment Charges, Late Payment Charges, Document Collection Fees, Agreement Fees, Default Charges or NACH Dishonour Charges along with applicable taxes, shall be paid to and collected by the platform- Finnovation Tech Solution Private Limited (GST NO- 29AACCF7831R1ZV) (“Platform”), which is a distinct and separate entity from the Lender. Lender is not liable to collect any such fee & charges and to pay any applicable taxes on the same. I have accepted that the acceptance of Loan under this Application Form places no obligation on the Lender to approve the Loan.", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", fontsize)
	pdf.CellFormat(0, 10, "SELF-DECLERATION AND UNDERTAKING:", "0", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `1. I hereby apply for the Loan facility from the Lender as specified above.`, "0", "L", false)
	pdf.MultiCell(0, 8, `2. I represent that the information and details provided in this Application Form and the documents submitted by me are true, correct and that I have not concealed any information.`, "0", "L", false)
	pdf.MultiCell(0, 8, `3. I have read and understood the fees and charges as may be applicable to the Loan that I may avail from time to time.`, "0", "L", false)
	pdf.MultiCell(0, 8, `4. I confirm that no insolvency proceedings or suits for recovery of outstanding dues have been initiated and / or are pending against me.`, "0", "L", false)
	pdf.MultiCell(0, 8, `5. I hereby authorize Lender to exchange or share information and details relating to this Application Form its group companies or any third party, as may be required or deemed fit, for the purpose of processing this loan application and/or related offerings or other products / services that I may apply for from time to time.`, "0", "L", false)
	pdf.MultiCell(0, 8, `6. I hereby consent to and authorize Lender to increase or decrease the credit limit assigned to me basis Lender's internal credit policy.`, "0", "L", false)
	pdf.MultiCell(0, 8, `7. By submitting this Application Form, I hereby expressly authorize Lender to send me communications regarding various financial products offered by or from Lender, its group companies and / or third parties through telephone calls / SMSs / emails / post etc. including but not limited to promotional communications. And confirm that I shall not challenge receipt of such communications as unsolicited communication, defined under TRAI Regulations on Unsolicited Commercial Communications under the Do Not Call Registry.`, "0", "L", false)
	pdf.MultiCell(0, 8, `8. I understand and acknowledge that Lender has the absolute discretion, without assigning any reasons to reject my application and that Lender is not liable to provide me a reason for the same.`, "0", "L", false)
	pdf.MultiCell(0, 8, `9. That Lender shall have the right to make disclosure of any information including my Aadhar number relating to me including personal information, details in relation to Loan, defaults, security, etc. to the Credit Information Bureau of India (CIBIL) and/or any other governmental/regulatory/statutory or private agency / entity, credit bureau, RBI, CKYCR, including publishing the name as part of wilful defaulter's list from time to time, as also use for KYC information verification, credit risk analysis, or for other related purposes.`, "0", "L", false)
	pdf.MultiCell(0, 8, `10. I agree and accept that Lender may in its sole discretion, by itself or through authorised persons, advocate, agencies, bureau, etc. verify any information given, check credit references, employment details and obtain credit reports to determine creditworthiness from time to time.`, "0", "L", false)
	pdf.MultiCell(0, 8, `11. That I have not taken any loan from any other bank/ finance company unless specifically declared by me.`, "0", "L", false)
	pdf.MultiCell(0, 8, `12. That the funds shall be used for the Purpose specified above and will not be used for any illegal, speculative or antisocial purpose.`, "0", "L", false)
	pdf.MultiCell(0, 8, `13. I have clearly understood and accepted the late payment and other default charges listed above.`, "0", "L", false)
	pdf.MultiCell(0, 8, `14. I hereby confirm that I contacted the Lender for my own personal requirement of personal loan and no representative of Lender has emphasized or induced me directly / indirectly to make this application for the Loan.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, `I hereby confirm having read and understood the Standard Terms and Conditions applicable to this Loan and`, "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 10, `I am signing this Application Form after understanding and acceptance of each term.`, "0", "L", false)

	pdf.AddPage()

	pdf.CellFormat(0, 10, "STANDARD TERMS", "0", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "The Borrower may apply for loans by submitting the Application Form(s) and Lender i.e Usha Financial Services Private Limited as specified in the Application Form (\"Lender\") may agree to grant such loan (a \"Loan\") that are or will be governed by these terms and conditions (\"Standard Terms\") read together with the Application Form(s), and Most Important Terms and Conditions (\"MITC\") as exchanged between the parties (together referred to as \"Transaction Document\").", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "1. Definitions:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "In these Standard Terms unless there is anything repugnant to the subject or context thereof, the expressions listed below, if applicable, shall have the following meanings:", "0", "L", false)

	pdf.MultiCell(0, 8, `i. "Access Code(s)" means any authentication mode as approved, specified by the Lender including without limitation combination of username and password.`, "0", "L", false)

	pdf.MultiCell(0, 8, `ii. "Account" means the bank account where the Loan disbursement is requested and more specifically provided under the Application Form;`, "0", "L", false)

	pdf.MultiCell(0, 8, `iii. "Application Form" means the loan application form submitted by the Borrower to the Lender for applying and availing of the Facility, together with all other information, particulars, clarifications and declarations, if any, furnished by the Borrower or any other persons from time to time in connection with the Facility;`, "0", "L", false)

	pdf.MultiCell(0, 8, `iv. "Availability Period" means the period of 11 (eleven) months and 15 (fifteen) days from the date of sanction of the Facility (or such extended date as may be approved by the Lender, in its sole discretion);`, "0", "L", false)

	pdf.MultiCell(0, 8, `v. "Business Day" means a day (other than a Saturday or a Sunday) on which banks are open for general business in Delhi;`, "0", "L", false)

	pdf.MultiCell(0, 8, `vi. "Due Date" means such date(s) on which any payment becomes due and payable under the terms of the Transaction Documents (or otherwise);`, "0", "L", false)

	pdf.MultiCell(0, 8, `vii. "Facility" means the Loan applied by the Borrower;`, "0", "L", false)

pdf.MultiCell(0, 8, `viii. "Increased Costs" means
	a) a reduction in the rate of return from the Loan(s) or on the Lender's overall capital (including as a result of any reduction in the rate of return on capital brought about by more capital being required to be allocated by the Lender)
	b) any additional or increased cost including provisioning as may be required under or as may be set out
	in RBI regulations or any other such regulations from time to time; or
	c) a reduction of any amount due and payable under the Transaction Documents;`, "0", "L", false)

	pdf.MultiCell(0, 8, `ix. "Sanctioning Authority" includes the Reserve Bank of India.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "2. Interpretation", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "Capitalized terms used in these Standard Terms but not defined herein, shall have the meaning ascribed to such terms under the Application Form.", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "3. Sanction and Disbursement", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "Capitalized terms used in these Standard Terms but not defined herein, shall have the meaning ascribed to such terms under the Application Form.", "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `3.1. The Lender may agree to grant the Facility to the Borrower on the basis of the information and representations provided in the Application Form and other Transaction Documents. The Lender shall not be obliged to provide any written acceptance to the request of the Borrower as contained in the Application Form or any other Transaction Documents and may do so orally or by disbursement of a Loan (or a part thereof) requested thereunder by the Borrower. Grant of the Facility and acceptance of the Borrower's request shall be at the absolute discretion of the Lender and the Lender shall not be required to notify any rejection of the Application Form to the Borrower`, "0", "L", false)

	pdf.MultiCell(0, 8, `3.2. The Borrower may request for disbursement only if (a) no Event of Default or potential event of default has occurred or is continuing, and (c) no material adverse event in the opinion of the Lender has occurred.`, "0", "L", false)

	pdf.MultiCell(0, 8, `3.3. The Borrower's request for the Facility under the Application Form shall be irrevocable unless rejected by the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `3.4. In the event that the Lender accepts the Borrower's request for the Facility and sanctions the Facility, the Lender may, subject to satisfaction of all the conditions precedents, disburse Loans into the Account and the Borrower confirms that the Loan shall be utilized only for the Purpose and subject to the terms under the Transaction Documents Any such disbursement made by the Lender into the Account (whether in the name of the Borrower or any third party) shall be a Loan. Lender shall not be responsible for any dispute between Borrower and any such third party.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "4. Interest and other charges", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "Capitalized terms used in these Standard Terms but not defined herein, shall have the meaning ascribed to such terms under the Application Form.", "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `4.1. The Loan(s) under the Facility shall carry interest at the rate specified in the Transaction Documents.`, "0", "L", false)

	pdf.MultiCell(0, 8, `4.2. In case of an Event of Default, Borrower will, to the extent permitted by law, be required to pay interest (before as well as after judgment / award) and such other charges plus GST (as applicable) mentioned in this Applicable Form on the overdue amount to the other party on demand, for the period from (and including) the original Due Date for payment to (but excluding) the date of actual payment, at the Default Rate.`, "0", "L", false)

	pdf.MultiCell(0, 8, `4.3. The Borrower acknowledge and agree that (i) the rates of interest specified in the Transaction Documents are reasonable and that they represent genuine pre-estimates of the loss expected to be incurred by the Lender in the event of non-payment of any monies by the Borrower; and (ii) the Rate of Interest payable by the Borrower shall be subject to change prospectively based on the monetary policies as may be changed by the Reserve Bank of India and other factors impacting the interest rates.`, "0", "L", false)

	pdf.MultiCell(0, 8, `4.4. In case the Borrower wishes to prepay the Loan in full or prepay any instalment under the Facility before the Due Date, such prepayment will be subject to Prepayment Charges as mentioned in the Summary of the Loan Terms and in the MITC, as applicable to such prepayment. It is clarified that no waiver in the Rate of Interest shall begiven to the Borrower for such prepayment and the Borrower shall continue to pay the instalments to the Lender in the manner specified by the Lender pursuant to Clause 5.5 hereunder.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "5. Payments:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "Capitalized terms used in these Standard Terms but not defined herein, shall have the meaning ascribed to such terms under the Application Form.", "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `5.1. The Borrower shall make each payment under the Transaction Documents on or before the respective Due Date. No Due Date shall exceed the Tenure of the Facility.`, "0", "L", false)

	pdf.MultiCell(0, 8, `5.2. If the respective Due Date is not a Business Day, then the Borrower agrees that the payment shall be made on the preceding Business Day.`, "0", "L", false)

	pdf.MultiCell(0, 8, `5.3. All payments shall be made in freely transferable funds without any set off, counter claim or any deduction. The Borrower shall not deduct/withhold any tax at source from the payment to be made to the Lender. In case wherein the borrower is required to deduct any taxes then prior information to be submitted to Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `5.4. Notwithstanding anything to the contrary, the Lender may, at any time, without assigning any reason, the undisbursed portion of the Facility and can also recall any or all portion of the disbursed Loan on demand in the event of occurrence of an Event of Default. Upon such recall, the Loan and other amounts stipulated by the Lender shall be payable forthwith.`, "0", "L", false)

	pdf.MultiCell(0, 8, `5.5. The Borrower shall make repayment of the principal amount under the Loan(s) in such proportion and periodicity as may be provided in the Transaction Documents or as communicated by the Lender from time to time.`, "0", "L", false)

	pdf.MultiCell(0, 8, `5.6. The Borrower may, subject to payment of Prepayment Charges, prepay the outstanding principal amounts of any of the Loan(s) in full or part before the Due Date.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "6. Representations and Warranties:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `6.1 The Borrower makes the representations and warranties set out in this Section 6.1 to the Lender, in reliance of which the Lender may grant the Facility:`, "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `i. The Borrower has the competence and has obtained all authorizations (which is in full force and effect) to enter into and perform under the Transaction Documents and to carry on its business and operations as it is being or is proposed to be conducted;`, "0", "L", false)

	pdf.MultiCell(0, 8, `ii. The Facility once granted by the Lender under Transaction Documents constitutes legal, valid and binding obligations of the Borrower enforceable in accordance with their respective terms;`, "0", "L", false)

	pdf.MultiCell(0, 8, `iii. The Borrower is in compliance with all laws (including laws relating to environment, social and labour, anti-corruption and anti-money laundering) applicable to the Parties;`, "0", "L", false)

	pdf.MultiCell(0, 8, `iv. The entry into, delivery and performance by the Borrower of, and the transactions contemplated by the Transaction Documents, do not and will not conflict: (a) with any law; (b) with the constitutional documents, if any, of the Borrower; or (c) with any document which is binding upon the Borrower or on any of its assets;`, "0", "L", false)

	pdf.MultiCell(0, 8, `v. The Borrower has a valid agreement, engagement or arrangement with the Portal which is currently subsisting, and the Portal has not blacklisted, delisted, suspend or otherwise terminated the arrangement;`, "0", "L", false)

	pdf.MultiCell(0, 8, `vi. Where the accounts are required to be audited under applicable law, the most recent audited accounts of the Borrower: (a) have been prepared in accordance with applicable accounting principles and practices generally accepted and consistently applied;(b) have been duly audited by the auditors in accordance with applicable laws (c) represent a true and fair view of its financial condition as at the date to which they were drawn up, and there has been no material adverse effect since the date on which those accounts were drawn up;`, "0", "L", false)

	pdf.MultiCell(0, 8, `vii. Except to the extent disclosed to the Lender, no litigation, arbitration, administrative or other proceedings are pending or threatened against the Borrower or its assets, which, if adversely determined, might have a material adverse effect;`, "0", "L", false)

	pdf.MultiCell(0, 8, `viii. All information communicated to or supplied by or on behalf of the Borrower to the Lender from time to time, are true and fair/ true, correct and complete in all respects as on the date on which it was communicated or supplied; and (b) Nothing has occurred since the date of communication or supply of any information to the Lender which renders such information untrue or misleading in any respect;`, "0", "L", false)

	pdf.MultiCell(0, 8, `ix. ix. The Borrower is not a specially designated national or otherwise sanctioned, under sanctions (and related laws) promulgated by any Sanctioning Authority; and`, "0", "L", false)

	pdf.MultiCell(0, 8, `x. The Borrower has not taken any action and no other steps have been taken or legal proceedings started by or against it in any court of law/ other authorities for its insolvency, bankruptcy, winding up, dissolution, administration or re- organization or for the appointment of a receiver, administrator, administrative receiver, trustee or similar officer of the Borrower or of any or all if its assets.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `6.2. Each of the representations set out in Section 6.1 shall be deemed to be repeated on each day during the Tenure of the Facility.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "7. Covenants:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `7.1. The Borrower shall provide all payment mandates including but not limited to electronic clearing service (ECS) or national automated clearing house (NACH) mandate, as and when demanded by the Lender, to the Lender or to any person nominated by the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.2. Borrower shall ensure that all payment mandates including electronic clearing service (ECS) or national automated clearing house (NACH) mandate, if any, provided to the Lender or to any person nominated by the Lender, are honoured at all times and such mandates are not altered or amended without prior permission of the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.3. Borrower shall permit and facilitate inspection and audit of Borrower's premises, books and statements of accounts and other documents as may be required by the Lender, for the purposes of the Facility.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.4. The Borrower shall maintain its existence and shall carry on its business and operations in compliance with all applicable laws (including laws relating to environment, social and labour, anti-corruption and anti- money laundering) applicable to the Parties and with due diligence and efficiency and in accordance with sound technical, financial and managerial standards and business practices.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.5. The Borrower shall ensure that the obligations under the Transaction Documents shall at least rank pari passu with all its unsecured and unsubordinated obligations.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.6. The Borrower shall, within 3 (three) Business Days of demand by the Lender, pay the amount of any Increased Costs incurred by the Lender as a result of (i) the introduction of or any change in (or in the interpretation, administration or application of ) any law or regulation; (ii) compliance with any law or regulation made before or after the date of relevant Loan (including any law or regulation concerning capital adequacy, prudential norms, liquidity, reserve assets or tax) or (ii) on account of factors beyond the control of the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.7. The Borrower agrees that it shall indemnify and hold harmless the Lender, to the fullest extent permitted by applicable law, for all losses and liabilities (including due to claims by a third party), incurred by the Lender as a result of any breach by him/her of any Transaction Documents.`, "0", "L", false)
	pdf.MultiCell(0, 8, `7.8. Borrower shall submit the details of financials, sales details, stock and book debts statement on such periodicity as required by the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.9. Borrower shall promptly provide upon request of the Lender any further document/information as may be required by the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.10. Borrower undertakes to ensure that all information provided in the Application Form(s) shall remain true at all times during the Tenure of the Facility.`, "0", "L", false)
	pdf.MultiCell(0, 8, `7.11. The Borrower shall utilize the Facility for the Purpose only .`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.12. The Borrower shall not:`, "0", "L", false)

	pdf.MultiCell(0, 8, `i. use the fund for purchase of gold in any form, (including primary gold, gold bullion, gold jewellery, gold coins, units of gold exchange traded funds (ETF) and units of gold mutual funds).`, "0", "L", false)
	pdf.MultiCell(0, 8, `ii. directly or indirectly: (a) use the Facility in any transaction with or for the purpose of financing the activities of, any person/country currently subject to any sanctions by Sanctioning Authority; and (b) take part in or financing any activity, production, use of, trade in, distribution of, or otherwise involved in any exclusion-list; and`, "0", "L", false)

	pdf.MultiCell(0, 8, `iii. directly or indirectly, make or offer any payment, gift or other advantage which is intended to, or does, influence or reward any person (whether or not they are in the public sector) for acting in breach of an expectation of good faith, impartiality or trust or otherwise performing their function improperly.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.13. The Borrower shall provide end use certificate in a form and manner satisfactory to the Lender.`, "0", "L", false)
	pdf.MultiCell(0, 8, `7.14. The Borrower shall promptly notify the Lender in writing upon occurrence of any breach of Covenant or representation or occurrence of Event of Default and the steps, if any, being taken to remedy it.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.15. The Borrower shall from time to time, if required by the Lender, provide additional security, in a form and manner satisfactory to the Lender.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.16. The Borrower shall promptly notify the Lender of any breach of any representations, warranties, covenants, undertakings or any other terms of these Standard Terms together with the steps taken to remedy it. Upon the breach being reported, the Lender may, without prejudice to any of its rights under law or contract, in its sole discretion recommend implementation of corrective measures to remedy such breach in a form, manner and time as may be necessary or desirable to the Lender.`, "0", "L", false)
	pdf.MultiCell(0, 8, `7.17. The Borrower shall ensure that throughout the Tenure of the Facility a valid agreement, engagement or arrangement with the Portal shall subsist and the Portal has not blacklisted/delisted/suspended or otherwise terminated such arrangement. This shall be considered an event of default for the purposes of the Loan.`, "0", "L", false)

	pdf.MultiCell(0, 8, `7.18. Notwithstanding any of the provisions of the Indian Contract Act, 1872 or any other applicable law or anything contained in the Transaction Documents, the amounts repaid by the Borrower shall be appropriated first towards cost, charges and expenses and other monies; secondly towards interest on cost charges and expenses and other monies; thirdly towards interest on the delayed payments; fourthly towards interest payable under the Transaction Documents and lastly towards repayment of any principal amounts.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "8. Event of Default:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `The Borrower shall be deemed to have committed default on the occurrence of, inter-alia but not limited to, any one or more of the following events (hereinafter referred to as "Event of Default"):`, "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `i. Default has occurred in the payment of any monies in respect of the Facility (whether at stated Due Date, by acceleration or otherwise) under the terms of the Transaction Documents.`, "0", "L", false)
	pdf.MultiCell(0, 8, `ii. Default (other than a payment default) has occurred in the performance of any covenant, condition, agreement or obligation on the part of the Borrower under the Transaction Documents and such default has continued for a period of 10 (ten) days after notice in writing thereof has been given to the Borrower or as the case may be, to such other person, by the Lender (except where the Lender is of the opinion that such default is incapable of remedy, in which event, no notice shall be required).`, "0", "L", false)
	pdf.MultiCell(0, 8, `iii. The Borrower has, or there is a reasonable apprehension that the Borrower would, voluntarily or involuntarily become the subject of proceedings under any bankruptcy or insolvency law, or is voluntarily or involuntarily dissolved, or declared insolvent, or if the Borrower has taken or suffered to be taken any action for its re- organization, liquidation or dissolution or if a receiver or liquidator or a similar official has been appointed or allowed to be appointed in respect of all or any part of the assets of the Borrower or if an attachment or distrait has been levied on the Borrower's assets or any part thereof or certificate proceedings have been taken or commenced for recovery of any dues from the Borrower or if one or more judgments or decrees have been rendered or entered against the Borrower and such judgments or decrees are not vacated, discharged or stayed for a period of 30(thirty) days, and such judgments or decrees involve in the aggregate, a liability which could have a material adverse effect.`, "0", "L", false)
	pdf.MultiCell(0, 8, `iv. Death of any Borrower, or the change of constitution of the Borrower without the consent of the Lender.`, "0", "L", false)
	pdf.MultiCell(0, 8, `v. Breach of any representation, warranty, declaration or confirmation made or deemed to be made under the Transaction Documents.`, "0", "L", false)
	pdf.MultiCell(0, 8, `vi. The Borrower is unable or has admitted in writing its inability to pay any of its indebtedness as and when they mature or become due.`, "0", "L", false)
	pdf.MultiCell(0, 8, `vii. If the Borrower ceases or threatens to cease to carry on any of its businesses in its current form which could have a material adverse effect.`, "0", "L", false)
	pdf.MultiCell(0, 8, `viii. The security, if any, for the Facility is in jeopardy or ceases to have effect.`, "0", "L", false)
	pdf.MultiCell(0, 8, `ix. It is or becomes unlawful for any Borrower to perform any of its obligations under the Transaction Documents.`, "0", "L", false)

	pdf.MultiCell(0, 8, `8.2. An event of default howsoever described (or any event which with the giving of notice, lapse of time, determination of materiality or fulfilment of any other applicable condition or any combination of the foregoing would constitute an event of default) occurs under any agreement or any indebtedness of the Borrower or becomes capable at such time of being declared, due and payable under such agreements before it would otherwise have been due and payable.`, "0", "L", false)

	pdf.MultiCell(0, 8, `8.3. One or more events, conditions or circumstances (including any change in law) shall occur or exist which could have a material adverse effect.`, "0", "L", false)

	pdf.MultiCell(0, 8, `8.4. The Borrower shall promptly notify the Lender in writing upon becoming aware of any default and the steps, if any, being taken to remedy it.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "9. Consequences of Event of Default:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `9.1. Upon occurrence of any Event of Default, the Lender shall be entitled at its absolute discretion to inter alia:`, "0", "L", false)

	pdf.MultiCell(0, 8, `i. Call upon the Borrower to pay forthwith the outstanding balance of the Facility together with interest and all sums payable by the Borrower to the Lender;`, "0", "L", false)

	pdf.MultiCell(0, 8, `ii. Call upon the Borrower to pay all claims, costs, losses and expenses that may be incurred by the Lender because of any act or default on the part of the Borrower with respect to the Facility and/or for the recovery of the outstanding dues (including legal/attorney fee) and/or on account of failure of the Borrower of any of the terms and conditions under the Transaction Documents;`, "0", "L", false)

	pdf.MultiCell(0, 8, `iii. Enforce any rights available to it under any law or contract.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.MultiCell(0, 10, "10. Miscellaneous:", "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, "10.1. Online Transactions:", "0", "L", false)

	pdf.MultiCell(0, 8, `i. For ease of operation of the Borrower, Borrower shall have the option of applying for further facilities provided by the Lender using online secure platforms as may be specified by the Lender (hereinafter referred to as "Online Facility").The Facility shall be extended to the Borrower subject to the Borrower complying with the Lender's credit parameters and submitting all documents/information as may be required by Lender in such form as may be specified by the Lender from time to time. Lender may in its sole discretion reject the application for the facility/loan by the Borrower.`, "0", "L", false)

	pdf.MultiCell(0, 8, `ii. It shall be the sole responsibility of the Borrower to ensure that the Access Codes are not compromised or shared with any unauthorized users.`, "0", "L", false)

	pdf.MultiCell(0, 8, `iii. The Borrower expressly agrees and acknowledges to have read and understood the terms applicable for usage of the Online Facility and be bound by such terms and conditions (as amended by the Lender from time to time)at all times during the tenure of such Facility.`, "0", "L", false)

	pdf.MultiCell(0, 8, `iv. The Lender shall have no obligation to verify the authenticity of any transaction/instruction received or purported to have been received from the Borrower through the Online Facility or purporting to have been sent by the Borrower other than by means of verification of the Access Codes.`, "0", "L", false)

	pdf.MultiCell(0, 8, `v. All the records of the Lender with respect to the online request for facility arising out of the use of the Online Facility shall be conclusive proof of the genuineness and accuracy of the transaction. While the Lender and its affiliates shall endeavour to carry out the instructions promptly, they shall not be responsible for anydelay in carrying on the instructions due to any reason whatsoever, including due to failure of operational systems or any requirement of law.`, "0", "L", false)

	pdf.MultiCell(0, 8, `vi. Borrower can check the availability of a pre-approved offer that may be made by the Lender through Online Facility. Any pre-approved offer by the Lender does not constitute grant of facility to the Borrower and shall be subject to the terms as may be specified by Lender from time to time.`, "0", "L", false)

	pdf.MultiCell(0, 8, `vii. Borrower acknowledges and accepts that the Lender may permit/allow anybody quoting the correct Access Codes and other details to conduct the type of operations which are permitted under the Online Facility.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `10.2. Notices`, "0", "L", false)
	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `Any notice or request to be given or made by a party to the other shall be in writing. Such notice or request shall be deemed to have been duly received by the party to whom it is addressed if it is given or made at the address specified below or at such other address as may be agreed from time to time:`, "0", "L", false)

	pdf.MultiCell(0, 8, `For the Lender: 330, Mezzanine Floor, Patparganj Industrial Area, Patparganj, Delhi 110092.`, "0", "L", false)

	pdf.MultiCell(0, 8, `For the Borrower: The address as stated in the Application Form. The Lender may also agree to act on the basis of request made via registered email ID of the Borrower.`, "0", "L", false)

	pdf.MultiCell(0, 8, `Provided however that a notice or communication to any Borrower shall be deemed to be a notice or communication to other Borrower(s)..`, "0", "L", false)
	pdf.MultiCell(0, 8, `10.3. Any and all claims and disputes arising out of or in connection with the Transaction Documents or its performance shall be subject to the laws of India and shall be settled by arbitration by a single arbitrator to be appointed by the Lender. The seat of arbitration shall be in Delhi. The arbitration shall be conducted under the provisions of the Arbitration and Conciliation Act, 1996 or any statutory modification or re- enactment thereof for the time being in force and the award of such arbitrator shall be final and binding upon the Borrower and the Lender. Subject to the foregoing, the courts of Delhi shall have the exclusive jurisdiction on all claims or disputes arising under any of the Transaction Document.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.4. The Lender shall, without prejudice to any of its specific rights under any other agreements with the Borrower, at its sole discretion, be at liberty to apply any other money, amounts, securities and other property of the Borrower (whether singly or jointly with another or others) in possession of the Lender or any of its subsidiary/ affiliate/ associate company in or towards payment of the dues under Facility granted under the Transaction Documents. The borrower can avail Online Facility on the Portal.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.5. The Lender shall be entitled at the sole risk and cost of the Borrower to engage one or more person(s) to collect the Borrower's dues and shall further be entitled to share such information, facts and figures pertaining to the Borrower as the Lender deems fit. The Lender may also delegate to such person(s) the right and authority to perform and execute all such acts, deeds, matters and things connected herewith, or incidental thereto, as the Lender may deems fit. The Borrower recognizes, accepts and consents to such delegation.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.6. The Lender shall have the right to disclose or publish any information regarding the Borrower or guarantor(s) (if any) and any information and documents that they might possess from time to time to:`, "0", "L", false)
	pdf.MultiCell(0, 8, `i. any of its branches or with other banks, financial institutions, Credit Information Bureau of India Limited, credit reference or rating agencies/bureaus or other individuals/entities either in response to their credit inquiries directed to the Lender or in the event of the Borrower not complying with any terms and conditions herein or otherwise;`, "0", "L", false)

	pdf.MultiCell(0, 8, `ii. the Reserve Bank of India and/or any other statutory authority or official of the Government of India or that of any other state.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.7. The Borrower shall not assign or transfer all or any of its rights, benefits or obligations under the Transaction Documents. The Lender may, at any time, assign or transfer all or any of its rights, benefits and obligations under the Transaction Documents.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.8. In case of default committed by the Borrower, the Lender may also disclose and publish the information about the Borrower and its default with the Lender in the public domain including through social media.`, "0", "L", false)
	pdf.MultiCell(0, 8, `10.9. The Borrower acknowledges that the Lender has or may have business and other transactions with third parties (including those who are in the business of manufacturing, supplying or otherwise dealing with any asset being financed by the proceeds of this Facility) and hereby waives any conflict of interest that it may have on such arrangement. Further, the Borrower acknowledges that any contract or arrangement between the Lender and such third parties are independent of these Standard Terms.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.10. The Lender reserves the right to amend the terms of these Standard Terms (except amendment to Rate of Interest) by intimating the same to the Borrower. Rate of Interest shall not be changed without prior consent of the Borrower.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.11 The Lender shall be entitled to set-off all monies, securities, deposits and other assets and properties belonging to the Borrowers and/or the Guarantor(s) in the possession of the Lender, whether in, or on any account of the Lender or otherwise, whether held singly or jointly by the Borrowers with others and may appropriate the same.`, "0", "L", false)

	pdf.MultiCell(0, 8, `10.12 GST: All the charges and fees mentioned in this Agreement shall be subject to applicable GST as per applicable GST laws. Any Borrower who is not registered under Goods and Service Tax Act (GST) shall not be eligible to claim input of GST paid on charges levied. In case the Borrower is registered under Goods and Service Tax Act and want to avail GST input, then please connect with the Platform.`, "0", "L", false)
	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `For the purpose of this clause, the term \'GST\' or \'GST law\' shall include the Central Goods and Services Tax (\'CGST\'), the State Goods and Services Tax (\'SGST\'), Integrated Goods and Services Tax (\'IGST\'), Union Territory Goods and Services Tax (\'UTGST\') and any other taxes levied under the GST related legislations in India as may be applicable. The term \'GST legislation/s\' should be accordingly interpreted.`, "0", "L", false)

	pdf.AddPage()

	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "ANNEXURE A", "0", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(0, 10, "THE MOST IMPORTANT TERMS AND CONDITONS - MITC", "0", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `1. We refer to the application form dated ______ ("Application Form") for grant of the Loan described below.`, "0", "L", false)

	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `2. Capitalized terms used but not defined hereunder shall have the meaning ascribed to the term in other Transaction Documents.`, "0", "L", false)

	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 8, `3. The Borrower acknowledges and confirms that the below mentioned are the most important terms and conditions in the application for the Loan (and which would apply to the Borrower in respect of the Loan, if the request for the Loan is accepted by the Lender) and they shall be read in conjunction with the Application Form(s), and the Standard Terms):`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Borrower", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Facility", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Facility Type", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Max Amount Sanctioned", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Purpose", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Tenure", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Rate of Interest (p.a.%)", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "EMI Amount Default Charges", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "NACH Dishonour Charges Repayment", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Prepayment Charges:", "LR", 0, "", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "The amount equivalent to the differential between the New Outstanding", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "(iii) In case of prepayment", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "Balance and the Existing Outstanding Balance shall be paid as", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "of any instalment/", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "prepayment charges by the Borrower in case of prepayment of any", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "part prepayment:", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "instalment/ part prepayment of the Loan.", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "(iv) In case of prepayment", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "[pre-determined rate to be included].", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "in full:", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "For the purpose of the above Clause:", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "\"Existing Outstanding Balance\" means the Outstanding Balance which", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "was remaining due to be paid under the Loan prior to prepayment of the", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "relevant instalment under the Loan, as reflected under the repayment", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "schedule provided by the Lender under Clause 5.5.", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "\"New Outstanding Balance\" means the Outstanding Balance remaining", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "due under the Loan pursuant to and post prepayment of any instalment", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "under the Loan.", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "\"utstanding Balance\" means collectively the principal, interest,:", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "compound interest, default charges/ interest, any other charges, dues and", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "monies payable, costs and expenses, reimbursable, as outstanding from", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "time to time and whether any of them are due or not in relation to the", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(130, 10, "Loan.", "LR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 8, "", "LR", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(halfWidth, 0, "", "1", 0, "", false, 0, "")
	pdf.CellFormat(halfWidth, 0, "", "1", 0, "", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(50, 10, "Processing Fee Risk Category", "1", 0, "L", false, 0, "")
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(130, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `4. The Borrower understands that the Lender has adopted risk-based pricing, which is arrived by taking into account, broad parameters like the customers financial and credit profile. Further, the Borrower acknowledges and confirms that the Lender shall have the discretion to change prospectively the rate of interest and other charges applicable to the Loan.`, "0", "L", false)

	pdf.MultiCell(0, 8, `5. The Borrower acknowledges and confirms having received a copy of each Transaction Document and agrees that this letter is a Transaction Document.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.MultiCell(0, 8, `ACKNOWELDEGEMENT: Lender acknowledges receipt of your Application Form together with the Standard Term. We will revert within 5 working days subject to furnishing the necessary documents to Lender's satisfaction. CC Details.`, "0", "L", false)


	err := pdf.OutputAndClose(pw)
	if err != nil {
		utilities.Println(err)
	}

	return pw.buf
}

func genPDFWithLOGO() []byte {
	pw := new(bufWriter)
	pdf := gofpdf.New("P", "mm", "A4", "")
	fontsize := 14.0
	smallfontsize := 10.0
	font := "Arial"
	// today := time.Now().Format("02/ 01/ 2006")
	logo := map[string]string{
		"imagePath":    "000000static%2Fdmi%20logo.png?alt=media&token=fb1d6202-27b5-48d4-bd9a-47b0a9beec3e",
		"imageKitLink": "https://firebasestorage.googleapis.com/v0/b/kreditbee-bucket/o/",
	}

	AddImageFromBuffer(logo["imagePath"], logo["imagePath"], logo["imageKitLink"], pdf, "")

	var options = gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	}
	pdf.SetHeaderFunc(func() {
		pdf.ImageOptions(logo["imagePath"], 5, 5, 70, 13, false, options, 0, "")
		pdf.Ln(-1)
		pdf.Ln(-1)
		pdf.Ln(-1)
	})

	pdf.AddPage()
	pdf.SetMargins(5, 5, 5)
	pdf.SetAutoPageBreak(true, 15)

	// addFooterToPDF(pdf, user, extraData)

	pdf.SetFont(font, "", fontsize)
	pdf.CellFormat(0, 25, "LOAN APPLICATION FORM", "0", 0, "C", false, 0, "")

	pdf.SetFont(font, "", smallfontsize)
	pdf.MoveTo(150, 20)
	pdf.CellFormat(25, 5, "Transaction ID: ", "1", 0, "L", false, 0, "")
	pdf.CellFormat(25, 5, "", "1", 0, "L", false, 0, "")
	pdf.MoveTo(150, 25)
	pdf.CellFormat(25, 5, "Date: ", "1", 0, "L", false, 0, "")
	pdf.CellFormat(25, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "CUSTOMER PERSONAL INFORMATION", "1", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(50, 5, "Full Name", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Mobile", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Date of Birth", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Personal E-mail", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Gender", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Landline Number", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Marital Status", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "PAN No", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(4)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "HOME INFORMATION", "1", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.CellFormat(50, 5, "Building Name & Apt #", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "Residence is", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "1", 0, "L", false, 0, "")

	pdf.Ln(5)
	pdf.CellFormat(50, 5, "Street Name", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Area", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Pincode Duration", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "City & State", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "1", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(50, 5, "in the City", "LRB", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "at Current Resi-dence", "LRB", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "City & State", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Duration in the City", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(4)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "WORK INFORMATION", "1", 0, "C", false, 0, "")

	pdf.SetFont(font, "", smallfontsize)
	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Type of Employment", "1", 0, "L", false, 0, "")
	pdf.CellFormat(150, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 10, "Company name", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "1", 0, "L", false, 0, "")

	pdf.CellFormat(50, 5, "Start Month / Year", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "1", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(50, 10, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "@ Current Job", "LRB", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Work Email", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Work Email Verified", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(4)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "SOURCING PARTNER", "1", 0, "C", false, 0, "")

	pdf.SetFont(font, "", smallfontsize)
	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Name", "1", 0, "L", false, 0, "")
	pdf.CellFormat(150, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Address", "1", 0, "L", false, 0, "")
	pdf.CellFormat(150, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Phone", "1", 0, "L", false, 0, "")
	pdf.CellFormat(150, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(4)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "UNDERTAKING", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.SetFont(font, "", smallfontsize)
	pdf.MultiCell(0, 10, `1. I/We hereby apply for a finance facility for the short term personal loan ("Facility") mentioned in this application. I/We declare that all the particulars and information and details given/filled in this Application Form are true, correct, complete and up-to date in all respects and no information has been withheld. I/We understand that the information given in this application shall form the basis of any loan that LENDER ("Lender") may decide to grant to me/us and if at any stage of processing this application, it comes to the knowledge of LENDER that, I/we have provided any incorrect or incomplete information, fabricated documents, or fake documents, they will be treated by LENDER as having been manipulated by me/us and LENDER shall have the right to forthwith reject this loan application, cancel / revoke any sanction or further drawdowns or recall any loan granted at any stage of processing the application, without assigning any reason whatsoever and LENDER and its employees/ representatives/ agents / service providers shall not be responsible/liable in any manner whatsoever to me/us for such rejection or any delay in notifying me/us of such rejetion (including for any payments which may have been made by me to any vendor/ service provider prior to cancellation).`, "0", "L", false)

	pdf.MultiCell(0, 10, `2. I/We understand that LENDER will also be procuring personal information from other sources/agents and I/We have no ob- jection for the same. I/We authorize LENDER to make reference and inquire relating to information in this application which LENDER considers necessary, including from the banks where I/we hold bank accounts.`, "0", "L", false)

	pdf.MultiCell(0, 10, `3. I/We hereby give my/our consent voluntarily for use of AADHAAR card/details for the purposes of availing the Facility in- cluding KYC purposes as required under applicable laws and as per LENDER policies. I/We hereby give our consent to LENDER to procure my /our AADHAAR details, PAN No/copy of my/our PAN Card, other identity proof and Bank Account details from time to time, exchange, part with/share all information relating to my/our loan details and repayment history with other banks/financial institutions /CIBIL etc. and periodically obtain / generate CIBIL, Experian, Hunter and such other reports as may be required and shall not hold LENDER liable for use of this information. I/We do hereby expressly and irrevocably au- thorize LENDER to collect, store, share, obtain and authenticate any aspect of my/our personal information either directly or through any of the authorized agencies and disclose such information to its agents/contractors/service providers and to also use such information for the purposes of KYC authentication, grant of the Facility and for internal evaluation by LENDER of its business.`, "0", "L", false)

	pdf.MultiCell(0, 10, `4. I/We, would like to know through telephonic calls, or SMS on my mobile number mentioned in the Application Form as well as in this undertaking, or through any other communication mode, transactional information, various loan offer schemes or loan promotional schemes or any other promotional schemes which may be provided by LENDER and hereby authorize LENDER and their employee, agent, associate to do so. I confirm that laws in relation to the unsolicited communication re- ferred in "National Do Not Call Registry" (the "NDNC Registry") as laid down by TELECOM REGULATORY AUTHORITY OF INDIA will not be applicable for such communication/calls/ SMSs received from the Lender, its employees, agents and/or as- sociates.`, "0", "L", false)

	pdf.MultiCell(0, 10, `5. I/We acknowledge that Sourcing Partner and LENDER are independent of each other and I/we will not have any claim against LENDER for any loan or other facility arranged/ provided by Sourcing Partner which is not sanctioned/ disbursed by Lender. I acknowledge that LENDER does not in any manner make any representation, promise, statement or endorsement in respect of any other product of services which may be provided by Sourcing Partner and will not be responsible or liable in any manner whatsoever for the same.`, "0", "L", false)

	pdf.MultiCell(0, 10, `6. I/we understand that the I/we have an option of not providing the information as required in this application form or as may be required by LENDER from time to time provided that on exercising such option LENDER shall have the right to cancel the sanction or seek prepayment of the amounts due as per the terms of the GC.`, "0", "L", false)

	pdf.MultiCell(0, 10, `7. I/ We declare that I/ We have not made any payment in cash, bearers cheques or by any other mode along with or in con- nection with this Application Form to the person collecting my/our Application Form. I/ We shall not hold LENDER or its em- ployees/representatives/agents/service providers liable for any such payment made by us to the person collecting this Ap- plication Form.`, "0", "L", false)

	pdf.Ln(-1)
	pdf.SetFont(font, "BU", 10)
	pdf.MultiCell(0, 10, `8. I/We confirm that I/We have understood the terms of the application form.`, "0", "L", false)

	pdf.SetFont(font, "", smallfontsize)
	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "Phone Number", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Signature", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, "DNC", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(4)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "ACKNOWLEDGMENT", "1", 0, "C", false, 0, "")

	pdf.SetFont(font, "", smallfontsize)
	pdf.Ln(-1)
	pdf.MultiCell(200, 5, `I/We acknowledge the receipt of your application for a loan from LENDER. Please furnish documents mentioned for approval of your loan application.`, "1", "L", false)

	pdf.Ln(0)
	pdf.CellFormat(50, 5, `Applicant's Name`, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Sales Representative Name", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `Loan Amount`, "LR", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Sales Representative", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "", "1", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(50, 5, "", "LRB", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "LRB", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Mobile No.", "LBR", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `Processing Fee: `, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Prepayment Charges", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `Interest Rate:`, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Bounce Charges:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.MultiCell(200, 5, `Overdue Interest Rate will be ___________ PM on an amount of EMI outstanding as on that date. Please note each application will be treated as different loan. Processing time of loan will be up to _________ days.`, "1", "L", false)
	pdf.CellFormat(120, 10, "", "LR", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "", "LR", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(120, 10, "Signature of the Borrower", "LRB", 0, "L", false, 0, "")
	pdf.CellFormat(120, 10, "Signature of Lender", "LRB", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.Ln(4)
	pdf.SetFont(font, "B", smallfontsize)
	pdf.CellFormat(0, 10, "DOCUMENT LIST", "1", 0, "C", false, 0, "")

	pdf.SetFont(font, "", smallfontsize)

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `Bank Statement (last 6 months)`, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "PDC / ECS Form", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `Pay Slip (last 3 months)`, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Cancelled Cheque", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `PAN Copy`, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Undated Cheque", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(50, 5, `Proof of Address`, "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "Aadhaar details", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, "", "1", 0, "L", false, 0, "")

	pdf.AddPage()

	fontsize = 10
	pdf.SetFont(font, "BI", fontsize)
	pdf.CellFormat(0, 10, "GENERAL TERMS AND CONDITIONS OF LOAN", "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `GENERAL TERMS AND CONDITIONS OF LOAN ("GC") for loans by the lender ('LENDER' which shall mean and include its successors and assigns):`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `1. DEFINITIONS`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `1.1 The terms and expressions contained in these GC and the Loan Application Form are defined as under:`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Availability Period" shall mean the period within which the Borrower can request a Drawdown from the Facility and is as detailed in the Loan Details Sheet;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Available Facility Amount" means at any point of time the undrawn amount of the Facility, including any amount of the Facility which becomes available pursuant to any repayment or prepayment of all or part of any previous Drawdown.`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Borrower" means the borrower as described in the Loan Details Sheet;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Borrower's Dues" means all sums payable by the Borrower to LENDER, including outstanding Facility, interest, all other charges, costs and expenses;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Drawdown" shall mean each drawdown of the Facility within the Availability Period and as per the terms of the Financing Documents, including drawdown of any amount which becomes available against the Facility, pursuant to prepayment/ repayment of any earlier Drawdown;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Due Date" in respect of any payment means the date on which any amount is due from the Borrower to LENDER;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"EMI" means the equated monthly amount to be paid by the Borrower towards repayment of all outstanding Drawdowns and payment of interest (if applicable) as per Financing Documents;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Facility" means the maximum drawdown limit granted by LENDER to the Borrower as per Loan Details Sheet, which may be available to the Borrower as a revolving credit;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Financing Documents" means these GC, the Loan Application, the Loan Details Sheet, including the annexures hereto and any documents executed by the Borrower or as required by LENDER, as amended from time to time;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Loan Application" means the application in the prescribed form as submitted from time to time by the Borrower to LENDER for seeking financing;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Loan Details Sheet" means the Loan Details Sheet executed between LENDER and Borrower, from time to time;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Material Adverse Effect" means any event which in LENDER's opinion would have an adverse effect on (i) Borrower's ability to pay the Borrower's Dues or (ii) recoverability of the Borrower's Dues;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Overdue Interest Rate" means the default interest as prescribed in the Loan Details Sheet which is payable on all amounts which are not paid on their respective Due Dates;`, "0", "L", false)
	pdf.MultiCell(0, 10, `"Purpose" means the utilization of each Drawdown as mentioned in the Loan Details Sheet.`, "0", "L", false)
	pdf.MultiCell(0, 10, `1.1A. In this GC, (a) the singular includes the plural (and vice versa) and (b) reference to a gender shall include references to the female,	male and neutral genders.`, "0", "L", false)

	pdf.Ln(-1)

	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `1. DISBURSEMENT`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `1.2 The Borrower may request for disbursement only if (a) no current or potential Material Adverse Effect has occurred, (b) Drawdown during the Availability Period for an amount not exceeding the Available Facility Amount has been submitted by the Borrower. LENDER shall have the sole and absolute discretion to allow or reject Drawdown against such request. The Facility may be in the nature of a revolving credit and the Available Facility Amount may change during the Availability Period on account of prepayments/repayment. Notwithstanding anything contained in this GC, LENDER shall have the absolute right to cancel or refuse any further Drawdowns from the Facility at its sole discretion as it may deem fit, including on account of any change in credit evaluation of the Borrower.`, "0", "L", false)
	pdf.MultiCell(0, 10, `1.3 The Borrower shall pay non-refundable processing charges as stated in the Loan Details Sheet, along with tax thereof, which may be added as a deemed disbursement to the first Drawdown and the Borrower will accordingly be liable for entire Drawdown.`, "0", "L", false)
	pdf.MultiCell(0, 10, `1.4 The Borrower's request for the Facility under this GC and each Drawdown shall be irrevocable unless rejected by the Lender.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `2. INTEREST AND REPAYMENT`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `2.1 The Borrower will pay Interest (if applicable) on each Drawdown made by the Borrower of the Facility and all other amounts due as provided in Loan Details Sheet and the interest shall be compounded on a monthly basis. The Borrower will be liable for the entire Drawdown amount and shall pay the full amount for each Drawdown. However, in such cases, in the event the installment is not paid on the Due Date, all overdue amounts shall accrue Interest at the prescribed rate ("Overdue Interest Rate") which shall be computed from the respective due dates for payments and the interest shall be compounded on a monthly basis. `, "0", "L", false)
	pdf.MultiCell(0, 10, `2.2 The tenure of each Drawdown shall be as provided in the Loan Details Sheet. EMI shall be as calculated by LENDER as required for amortization of Drawdowns within their respective tenure and Interest payable thereon and not exceeding the maximum EMI as provided in the Loan Details Sheet. EMI shall only be towards principal outstanding and Interest thereon and does not include any default interest or any other charges payable by the Borrower pursuant to Financing Documents.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.3 The payment of each EMI on time is the essence of the contract. The Borrower acknowledges that s/he has understood the method of computation of EMI and shall not dispute the same.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.4 Notwithstanding anything stated elsewhere in the Financing Documents, all Borrower's Dues, including EMI, shall be payable by the Borrower to LENDER as and when demanded by LENDER, at its sole discretion and without requirement of any reason being assigned. The Borrower shall pay such amounts, without any delay or demur, within 15 (fifteen) days of such demand.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.5 LENDER shall be entitled to revise the rate of interest, if so required under any applicable law and LENDER may recompute the EMI /the number of EMI for repayment of outstanding Facility and interest. Any such change as intimated by LENDER to Borrower will be final and binding on the Borrower. In case of such revision the Borrower shall be entitled to prepay, within 30 (thirty) days of such revision, the entire outstanding Facility along with accrued Interest (if applicable), without any prepayment penalty.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.6 In case of delayed payments, without prejudice to all other rights of LENDER, LENDER shall be entitled to Overdue Interest Rate (as prescribed in Loan Details Sheet) from the Borrower for the period of delay.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.7 The Borrower may pre-pay any Drawdown prior to its scheduled tenure only with the prior approval of LENDER and subject to such conditions and prepayment charges, as stipulated by LENDER as mentioned in the Loan Details Sheet. It is clarified that Interest on the Drawdown being prepaid, shall be charged only for the period preceding the date of prepayment by the Borrower on pro rata basis, and not for the unexpired Tenure of the Drawdown. The Borrower may pre-pay any Drawdown prior to its scheduled tenure only with the prior approval of LENDER and subject to such conditions and prepayment charges, as stipulated by LENDER.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.8 The Borrower shall bear all interest, tax, duties, cess duties and other forms of taxes whether applicable now or in the future, payable`, "0", "L", false)
	pdf.MultiCell(0, 10, `under any law at any time in respect of any payments made to LENDER under the Financing Documents. If these are incurred by LENDER, these shall be recoverable from the Borrower and will carry interest at the rate of Overdue Interest Rate from the date of payment till reimbursement.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.9 Notwithstanding any terms and conditions to the contrary contained in the Financing Documents, the amounts repaid by the Borrower shall be appropriated firstly towards cost, charges, expenses and other monies; secondly towards Overdue Interest Rate, if any; thirdly towards Interest; and lastly towards repayment of principal amount of a Facility.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.10 Interest (if applicable), Overdue Interest Rate and all other charges shall accrue from day to day and shall be computed on the basis of 365 days a year and the actual number of days elapsed.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.11 If the due date for any payment is not a business day, the amount will be paid by Borrower on immediately succeeding business day.`, "0", "L", false)
	pdf.MultiCell(0, 10, `2.12 All sums payable by the Borrower to LENDER shall be paid without any deductions whatsoever. Credit/ discharge for payment will be given only on realization of amounts due.`, "0", "L", false)

	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `3. MODE OF PAYMENT, REPAYMENT AND PREPAYMENT`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `3.1 The Borrower shall, as required by LENDER from time to time, provide National Automated Clearing House (Debit Clearing)/ any other electronic or other clearing mandate (collectively referred to as "NACH") as notified by the Reserve Bank of India ("RBI") against Borrower's bank account for payment of dues. Such NACH shall be drawn from such bank and from such location as agreed to by LENDER. The Borrower shall honor all payments without fail on first presentation/ due dates. NACH provided by the Borrower may be utilized by LENDER for realization of any Borrower's Dues. The Borrower hereby unconditionally and irrevocably authorizes LENDER to take all actions required for such realization. The Borrower shall promptly (and in any event within seven (5) days) replace the NACH and/or other documents executed for payment of Borrower's Dues as may be required by LENDER from time to time, at its sole discretion.`, "0", "L", false)
	pdf.MultiCell(0, 10, `3.2 The Borrower shall, at all times maintain sufficient funds in his/her bank account/s for due payment of the Borrower's Dues on respective Due Dates. Borrower shall not close the bank account/s from which the NACH have been issued or cancel or issues instructions to the bank or to LENDER to stop or delay payment under the NACH and LENDER is not bound to take notice of any such communication.`, "0", "L", false)
	pdf.MultiCell(0, 10, `3.3 The Borrower agrees and acknowledges that NACH have been issued voluntarily in discharge of the Borrower's Dues and not by way of a security for any purpose whatsoever. The Borrower also acknowledges that dishonor of any NACH is a criminal offence under the Negotiable Instruments Act, 1881/The Payment and Settlements Act, 2007. The Borrower shall be liable to pay dishonour charges for each NACH dishonour (as prescribed in Loan Details Sheet).`, "0", "L", false)
	pdf.MultiCell(0, 10, `3.4 Any dispute or difference of any nature whatsoever shall not entitle the Borrower to withhold or delay payment of any EMIs or other sum and LENDER shall be entitled to present NACH on the respective due dates.`, "0", "L", false)
	pdf.MultiCell(0, 10, `3.5 Notwithstanding the issuance of NACH, the Borrower will be solely responsible to ensure timely payment of dues.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `4. BORROWER'S COVENANTS, REPRESENTATION AND WARRANTIES`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `4.1 The Borrower shall:`, "0", "L", false)
	pdf.MultiCell(0, 10, `(i) observe and perform all its obligations under the Financing Documents.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(ii) immediately deliver to LENDER all documents, including bank account statements as may be required by LENDER from time to time. The Borrower also authorizes LENDER to communicate independently with (i) any bank where the Borrower maintains an account and to seek details and statement in respect of such account from the bank and (ii) with any employer of any Borrower as LENDER may deem necessary, including for monitoring Borrower's creditworthiness.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(iii) immediately notify LENDER of any litigations or legal proceedings against any Borrower.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(iv) notify LENDER of any Material Adverse Effect or Event of Default.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(v) notify LENDER in writing of all changes in the location/ address of office /residence /place of business or any change/resignation/termination / closure of employment/ profession /business.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(vi) Not leave India for employment or business or long term stay abroad without fully repaying the Facility then outstanding, together with interest and other dues and charges.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(vii) provide security, if any, as specified in Financing Documents or as may be required by LENDER in case of any change in credit worthiness of any Borrower (as determined by LENDER).`, "0", "L", false)
	pdf.MultiCell(0, 10, `(viii) Ensure deposit of salary and / or business proceeds in the account from which PDCs/ECS have been issued to LENDER.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(ix) On or prior to the first Drawdown take a credit life insurance policy as required by LENDER which shall include a cover for accidents, death, permanent disability and unemployment and such other terms as shall be acceptable to LENDER.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(x) comply at all times with applicable laws, including, Prevention of Money Laundering Act, 2002.`, "0", "L", false)
	pdf.MultiCell(0, 10, `(xi) Utilise each Drawdown only for the Purpose.`, "0", "L", false)
	pdf.MultiCell(0, 10, `4.2 Each Borrower represents and warrants to LENDER as under:`, "0", "L", false)
	pdf.MultiCell(0, 10, `(i) All the information provided by Borrower in the Loan Application and any other document, whether or not relevant for the ascertaining the credit worthiness of the Borrower, is true and correct and not misleading in any manner;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(ii) The Borrower is capable of and entitled under all applicable laws to execute and perform the Financing Documents and the transactions thereunder;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(iii) The Borrower is above 18 years of age and this GC is a legal, valid and binding obligation on him/her, enforceable against him/her in accordance with its terms;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(iv) The Borrower declares that he/she is not prohibited by any law from availing this Facility;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(v) No event has occurred which shall prejudicially affect the interest of LENDER or affect the financial conditions of Borrower or affect his/her liability to perform all or any of their obligations under the Financing Documents;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(vi) Borrower is not in default of payment of any taxes or government dues;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(vii) The Borrower will do all acts, deeds and things, as required by LENDER to give effect to the terms of this GC;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(viii) Commencement of any bankruptcy or insolvency proceedings against the Borrower. The Borrower gives its consent to LENDER to use/store all the information provided by the Borrower or otherwise procured by LENDER in the manner it deems fit including for the purposes of this Facility or for its business and understands and agrees that LENDER may disclose such information to its contractors, agents and any other third parties.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `5. EVENTS OF DEFAULT`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `5.1 The following acts/events, shall each constitute an "Event of Default" by the Borrower for the purposes of each Facility:`, "0", "L", false)
	pdf.MultiCell(0, 10, `(i) The Borrower fails to make payment of any Borrower's Dues on Due Date;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(ii) Breach of any terms, covenants, representation, warranty, declaration or confirmation under the Financing Documents;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(iii) Any fraud or misrepresentation or concealment of material information by Borrower which could have affected decision of LENDER to grant any Facility;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(iv) Death, lunacy or any other permanent disability of the Borrower;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(v) Borrower utilises the Drawdown for any purpose other than the Purpose;`, "0", "L", false)
	pdf.MultiCell(0, 10, `(vi) Occurrence of any events, conditions or circumstances (including any change in law) which in the sole and absolute opinion of LENDER could have a Material Adverse Effect, including limitation of any proceedings or action for bankruptcy/liquidation/ insolvency of the Borrower or attachment / restraint of any of its assets;`, "0", "L", false)
	pdf.MultiCell(0, 10, `5.2 The decision of LENDER as to whether or not an Event of Default has occurred shall be binding upon the Borrower.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)

	pdf.MultiCell(0, 10, `6. CONSEQUENCES OF DEFAULT`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `6.1 Upon occurrence of any of the Events of Default and at any time thereafter, LENDER shall have the right, but not the obligation to declare all sums outstanding in respect of the Facility, whether due or not, immediately repayable and upon the Borrower failing to make the said payments within 15 (fifteen) days thereof, LENDER may at its sole discretion exercise any other right or remedy which may be available to LENDER under any applicable law, including seeking any injunctive relief or attachment against the Borrower or their assets.`, "0", "L", false)
	pdf.MultiCell(0, 10, `6.2 The Borrower shall also be liable for payment of all legal and other costs and expenses resulting from the foregoing defaults or the exercise of LENDER remedies.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `7. DISCLOSURES`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `7.1 The Borrower acknowledges and authorizes LENDER to disclose all information and data relating to Borrower, the Facility, Drawdowns, default if any, committed by Borrower to such third parties/ agencies as LENDER may deem appropriate and necessary to disclose and/or as authorized by RBI, including the TransUnion CIBIL Limited (CIBIL). The Borrower also acknowledges and authorizes such information to be used, processed by LENDER / third parties/ CIBIL / RBI as they may deem fit and in accordance with applicable laws. Further in Event of Default, LENDER and such agencies shall have an unqualified right to disclose or publish the name of the Borrower /or its directors/ partners/co-applicants, as applicable, as ‘defaulters' in such manner and through such medium`, "0", "L", false)
	pdf.MultiCell(0, 10, `as LENDER / CIBIL/ RBI/ other authorized agency in their absolute discretion may think fit, including in newspapers, magazines and social media.`, "0", "L", false)
	pdf.MultiCell(0, 10, `7.2 The Borrower shall not hold LENDER responsible for sharing and/or disclosing the information now or in future and also for any consequences suffered by the Borrower and/or other by reason thereof. The provisions of this clause 8 shall survive termination of the GC and the repayment of the Borrower's Dues.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `8. MISCELLANEOUS`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `8.1 The entries made in records of LENDER shall be conclusive evidence of existence and of the amount Borrower's Dues and any statement of dues furnished by LENDER shall be accepted by and be binding on the Borrower.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.2 Borrower's liability for repayment of the Borrower's Dues shall, in case where more than one Borrower have jointly applied for any Facility, be joint and several.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.3 Borrower shall execute all documents and amendments and shall co-operate with LENDER as required by LENDER (i) to comply with any RBI guidelines / directives or (ii) for giving LENDER full benefit of rights under the Financing Documents. Without prejudice to the aforesaid the Borrower hereby irrevocably consents that on its failure to do so, such changes shall be deemed to be incorporated in the Financing Documents and shall be binding on the Borrower.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.4 Notwithstanding any suspension or termination of any Facility, all right and remedies of LENDER as per Financing Documents shall continue to survive until the receipt by LENDER of the Borrower's Dues in full.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.5 The Borrower acknowledges that the rate of interest, penal charges, service charges and other charges payable and or agreed to be paid by the Borrower under Financing Documents are reasonable and acceptable to him/ her.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.6 The Borrower expressly recognizes and accepts that LENDER shall, without prejudice to its rights to perform such activities itself or through its office employees, be entitled and has full power and authority so to appoint one or more third parties (hereinafter referred to as "Service Providers") as LENDER may select and to delegate to such party all or any of its functions, rights and power under Financing Documents relating to the sourcing, identity and verification of information pertaining to the Borrower administration, monitoring of the Facility and to perform and execute all lawful acts, deeds, matters and things connected therewith and incidental thereto including sending notices, contacting Borrower, receiving Cash / Cheques/ Drafts / Mandates from the Borrower in`, "0", "L", false)
	pdf.MultiCell(0, 10, `favour of LENDER.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.7 The Borrower acknowledges that the financing transaction hereunder gives rise to a relationship of debtor and creditor as between him / her and LENDER and not in respect of any service rendered/to be rendered by LENDER. Accordingly, the provisions of the Consumer Protection Act, 1986 shall not apply to the transaction hereunder.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.8 The Borrower hereby authorizes LENDER to verify all information and documents including, income proof documents, residence documents, address proof documents, identity documents and other such documents containing personal and financial information as are submitted by them for obtaining any Facility and that they also consent to subsequent retention of the same by LENDER.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.9 The Borrower acknowledges and authorizes LENDER to procure Borrower's PAN No./copy of Pan Card, other identity proof and Bank Account details, from time to time and to also generate / obtain CIBIL, Experian, Hunter reports and such other reports as and when LENDER may deem fit. The Borrower also hereby gives consent and authorizes LENDER to undertake its KYC verification by Aadhar e-KYC or otherwise and undertake all such actions as may be required on its behalf or otherwise to duly complete the process of such verification including by way of Aadhar e-KYC and share such information with any authority and store such information in a manner it deems fit.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.10 In the event of any disagreement or dispute between LENDER and the Borrower regarding the materiality of any matter including of any event occurrence, circumstance, change, fact information, document, authorization, proceeding, act, omission, claims, breach, default or otherwise, the opinion of LENDER as to the materiality of any of the foregoing shall be final and binding on the Borrower.`, "0", "L", false)
	pdf.MultiCell(0, 10, `8.11 The Borrower and LENDER may mutually agree on grant of a fresh facility on the terms and conditions of the GC and by execution of such further letter/undertaking by the Borrower as may be required by LENDER.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `9. SEVERABILITY `, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `The Borrower acknowledges that each of his /her obligations under these Financing Documents is independent and severable from therest.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `10. GOVERNING LAW AND JURISDICTION`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `10.1 All Facility and the Financing Documents shall be governed by and construed in accordance with the laws of India.`, "0", "L", false)
	pdf.MultiCell(0, 10, `10.2 All disputes, differences and / or claims arising out of these presents or as to the construction, meaning or effect hereof or as to the right and liabilities of the parties under the Financing Documents shall be settled by arbitration in accordance with the provision of the Arbitration and Conciliation Act, 1996 or any statutory amendments thereof or any statute enacted for replacement therefore and shall be referred to the sole Arbitration of a person to be appointed by LENDER. The place of arbitration shall be Delhi and proceeding shall be under fast track procedure as laid down in Section 29(B) of the Act. The awards including interim awards of the arbitration shall be final and binding on all parties concerned. The arbitrator may pass the award without stating any reasons in such award.`, "0", "L", false)
	pdf.MultiCell(0, 10, `10.3 Further, the present clause shall survive the termination of Financing Documents. The Courts at Delhi, India shall have exclusive jurisdiction (subject to the arbitration proceedings which are to be also conducted in Delhi, India) over any or all disputes arising out of the Financing Documents.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `11. NOTICES`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `11.1 Any notice to be given to the Borrower in respect of Financing Documents shall be deemed to have been validly given if served on the Borrower or sent by registered post to or left at the address of the Borrower existing or last known business or private address. Any such notice sent by registered post shall be deemed to have been received by the Borrower within 48 hours from the time of its posting. Any notice to LENDER shall be deemed to have been valid only if received by LENDER at its abovestated address.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `12. ASSIGNMENT`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `12.1 The Borrower shall not be entitled to jointly or severally transfer or assign all or any of their right or obligation or duties under the Financing Documents to any person directly or indirectly or create any third party interest in favour of any person without the prior written consent of LENDER.`, "0", "L", false)
	pdf.MultiCell(0, 10, `12.2 LENDER shall be entitled to sell, transfer, assign or securitise in any manner whatsoever (in whole or in part and including through grant of participation rights) all or any of its benefits, right, obligation, duties and / or liabilities under Financing Documents, without the prior written consent of, or intimation to the Borrower in such manner and such terms as LENDER may decide. In the event of such transfer, assignment or securitization, the Borrower shall perform and be liable to perform their obligation under the Financing Documents to such assignee or transferor. In such event, the Borrower shall substitute the remaining PDCs/ECS in favour of the transferee/ assignee if called upon to do so by LENDER.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `13. INDEMNITY`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `The Borrower hereby indemnifies, defends and holds LENDER, its employees, representatives and consultants harmless from time to time and at all times against any liability, claim, loss, judgment, damage, cost or expense (including, without limitation, reasonable attorney's fees and expenses) as a result of or arising out of any failure by the Borrower to observe or perform any of the terms and conditions and obligations contained in the Financing Documents or Event of Default or the exercise of any of the rights by LENDER under the Financing Documents, including for any enforcement of security or recovery of Borrower's Dues.`, "0", "L", false)
	pdf.SetFont(font, "BI", fontsize)
	pdf.MultiCell(0, 10, `14. Acceptance:`, "0", "L", false)
	pdf.SetFont(font, "I", fontsize)
	pdf.MultiCell(0, 10, `I / We am / are aware that LENDER shall agree to become a party to this GC only after satisfying itself with regard to all conditions and details filled by me / us in the GC and other Financing Documents in consonance with LENDER policy. I / We agree that this GC shall be concluded and become legally binding on the date when the authorized officer of LENDER signing this at Delhi or on the date of first disbursement, whichever is earlier. By clicking "I accept", the Borrower electronically signs these GC and agrees to be legally bound by its terms. The Borrower's acceptance of these GC shall constitute: (i) the Borrower's agreement to irrevocably accept and to be unconditionally bound by all the terms and conditions set out in these GC; and (ii) the Borrower's acknowledgement and confirmation that these GC (along with the Financing Documents) have been duly read and fully understood by the Borrower.`, "0", "L", false)

	err := pdf.OutputAndClose(pw)
	if err != nil {
		utilities.Println(err)
	}

	return pw.buf
}

// AddImageFromBuffer ... function  that add image from the buffer to the pdf
func AddImageFromBuffer(name, url string, baseURL string, f *gofpdf.Fpdf, imgTrans string) (string, string, []byte) {
	// mode := viper.GetString("providermode")
	// firebaseConf, err := config.GetConfigByType("firebaseStorage/" + mode)
	// bucket := firebaseConf["bucket"]
	// imgTrans := "tr:w-600,h-400,c-at_max/" // passing as parameter for image resize
	fullURL := baseURL + imgTrans + strings.TrimSpace(url)
	// log.Println("image link", fullURL)

	temp1 := strings.Split(url, "?")
	filename := strings.Split(temp1[0], ".")
	// filename[1] = "png"

	slurp := DownloadDocs(fullURL)

	// log.Println("image size", len(slurp))
	// log.Println("reading body", err)
	var options = gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: filename[1],
	}
	// log.Println(options)
	// log.Println(name, resp.Body)
	// utilities.LogErrorInRaven(err)

	if f != nil {
		f.RegisterImageOptionsReader(name, options, bytes.NewBuffer(slurp))
	}
	// log.Println("image options", check)
	return name, fullURL, slurp

}

func DownloadDocs(url string) (doc []byte) {
	var err error
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			break
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			break
		}

		doc, err = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		break
	}

	if err != nil {

	}

	return
}
