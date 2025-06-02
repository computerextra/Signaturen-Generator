package template

import (
	"Signaturen-Generator/db"
	"fmt"
	"os"
)

func GenerateTextFile(employee db.Employee, filePath string) error {
	fileName := fmt.Sprintf("%s/%s (%s).txt", filePath, employee.Name, employee.Mail)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("Mit freundlichen Grüßen\n\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString(employee.Name)
	if err != nil {
		return err
	}
	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString(employee.Abteilung)
	if err != nil {
		return err
	}
	_, err = file.WriteString("\n\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("Telefon: 0561 60 144 ...\n")
	if err != nil {
		return err
	}
	fax := os.Getenv("COMPANY_FAX")
	if len(fax) > 0 {
		_, err = file.WriteString(fmt.Sprintf("Fax: %s\n"))
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString(fmt.Sprintf("E-Mail: %s\n", employee.Mail))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Webseite: www.%s\n\n", os.Getenv("COMPANY_WEBSITE")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", os.Getenv("COMPANY_NAME")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", os.Getenv("COMPANY_STREET")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("%s %s\n\n", os.Getenv("COMPANY_POSTAL"), os.Getenv("COMPANY_CITY")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Sitz der Gesellschaft: %s %s\n", os.Getenv("COMPANY_POSTAL"), os.Getenv("COMPANY_CITY")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Geschäftsführer: %s – %s\n", os.Getenv("COMPANY_CHEF"), os.Getenv("COMPANY_REGISTER")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("USt.-IdNr.: %s\n\n", os.Getenv("COMPANY_VAT_ID")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Datenschutzinformationen: %s\n", os.Getenv("COMPANY_DATA_PROT_URL")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("AGB: %s\n", os.Getenv("COMPANY_GTC")))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Impressum: %s\n", os.Getenv("COMPANY_LEGAL")))
	if err != nil {
		return err
	}
	appointment := os.Getenv("COMPANY_APPOINTMENT_URL")
	if len(appointment) > 0 {
		_, err = file.WriteString(fmt.Sprintf("Vereinbaren Sie einen Termin Online: %s\n\n", appointment))
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("Der Inhalt dieser E-Mail und sämtliche Anhänge sind vertraulich und ausschließlich für den bezeichneten Empfänger bestimmt.\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("Sollten Sie nicht der bezeichnete Empfänger sein, bitten wir Sie, umgehend den Absender zu benachrichtigen und diese E-Mail zu löschen.\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("Jede Form der unautorisierten Veröffentlichung, Vervielfältigung und Weitergabe des Inhalts dieser E-Mail oder auch das Ergreifen von\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("Maßnahmen als Reaktion darauf sind unzulässig.\n")
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
