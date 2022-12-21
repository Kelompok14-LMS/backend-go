package certificates

import (
	"bytes"
	"html/template"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/templates"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type certificateUsecase struct {
	menteeRepository mentees.Repository
	courseRepository courses.Repository
}

func NewCertificateUsecase(menteeRepository mentees.Repository, courseRepository courses.Repository) Usecase {
	return certificateUsecase{
		menteeRepository: menteeRepository,
		courseRepository: courseRepository,
	}
}

func (cu certificateUsecase) GenerateCert(data *Domain) ([]byte, error) {
	// NOTE: in windows, should be point to wkhtmltopdf executable file (binary of wkhtmltopdf)
	// wkhtmltopdf.SetPath("C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf.exe")

	mentee, err := cu.menteeRepository.FindById(data.MenteeId)

	if err != nil {
		return nil, err
	}

	course, err := cu.courseRepository.FindById(data.CourseId)

	if err != nil {
		return nil, err
	}

	tmpl := template.Must(template.ParseFS(templates.Certificate, "template-certificate.html"))

	// prepare data certificate needs
	certDomain := map[string]string{
		"fullname": mentee.Fullname,
		"title":    course.Title,
	}

	// apply parsed HTML template data and keep the result in a buffer
	var w bytes.Buffer

	if err := tmpl.Execute(&w, certDomain); err != nil {
		return nil, err
	}

	// init a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewBuffer(w.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	// add the page to generator
	pdfg.AddPage(page)

	// manipulate attribute
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA5)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Cover.Zoom.Set(1.2)

	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
