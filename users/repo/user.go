/*package sql

import (
	"database/sql"
	"prdt-trnf-notifications-reports-operations/internal/core/domain"
	"prdt-trnf-notifications-reports-operations/internal/core/ports"
	"prdt-trnf-notifications-reports-operations/internal/utils/database"
	"prdt-trnf-notifications-reports-operations/internal/utils/env"
	"prdt-trnf-notifications-reports-operations/internal/utils/helper"

	log "github.com/sirupsen/logrus"
)

type reportClRepository struct {
	db *sql.DB
}

func NewReportClRepository(db *sql.DB) ports.ReportClRepository {
	return &reportClRepository{db: db}
}

func (r reportClRepository) GetCl(yesterdayDate string) ([][]string, error) {
	//fmt.Println("entra en la BD")
	//define schemaDB attacth
	if env.Enviroment == "dev" || env.Enviroment == "staging" {
		database.SetterSchema(r.db, env.SchemaDB)
	} else {
		database.SetterSchema(r.db, env.SchemaDBCl)
	}

	var reportData [][]string

	rows, err := r.db.Query(env.GetReportCl, yesterdayDate)

	if err != nil {
		log.Printf("cannot execute select query %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var reportClQuery domain.ReportOperationsCl

		err := rows.Scan(
			&reportClQuery.Recipt,
			&reportClQuery.Subject,
			&reportClQuery.OperationNumber,
			&reportClQuery.LineNumber,
			&reportClQuery.Product,
			&reportClQuery.Subproduct,
			&reportClQuery.Plan,
			&reportClQuery.TemplateId,
			&reportClQuery.Categories,
			&reportClQuery.Source,
			&reportClQuery.Flow,
			&reportClQuery.Country,
			&reportClQuery.Channel,
			&reportClQuery.CustomerId,
			&reportClQuery.NotificationType,
			&reportClQuery.ClientNotify,
			&reportClQuery.DateInitSent,
			&reportClQuery.ShippingStatus,
			&reportClQuery.DatetimeSent,
			&reportClQuery.ReceivedStatus,
			&reportClQuery.OpeningDateTime,
			&reportClQuery.BounceType,
			&reportClQuery.BounceMotive,
			&reportClQuery.BounceDateTime,
		)
		if err != nil {
			log.Fatal(err)
		}

		reportData = append(reportData, []string{
			reportClQuery.Recipt.String,
			reportClQuery.Subject.String,
			reportClQuery.OperationNumber.String,
			reportClQuery.LineNumber.String,
			reportClQuery.Product.String,
			reportClQuery.Subproduct.String,
			reportClQuery.Plan.String,
			reportClQuery.TemplateId.String,
			reportClQuery.Categories.String,
			reportClQuery.Source.String,
			reportClQuery.Flow.String,
			reportClQuery.Country.String,
			reportClQuery.Channel.String,
			reportClQuery.CustomerId.String,
			reportClQuery.NotificationType.String,
			reportClQuery.ClientNotify.String,
			helper.DateToString(reportClQuery.DateInitSent),
			reportClQuery.ShippingStatus.String,
			helper.DateToString(reportClQuery.DatetimeSent),
			helper.DateToString(reportClQuery.ReceivedStatus),
			helper.DateToString(reportClQuery.OpeningDateTime),
			reportClQuery.BounceType.String,
			reportClQuery.BounceMotive.String,
			helper.DateToString(reportClQuery.BounceDateTime),
		})
		//fmt.Println("deberia formatear")
		//fmt.Println(reportClQuery)
	}
	return reportData, nil
}

func (r reportClRepository) GetMassiveCl(yesterdayDate string) ([][]string, error) {
	//define schemaDB attacth
	if env.Enviroment == "dev" || env.Enviroment == "staging" {
		database.SetterSchema(r.db, env.SchemaDB)
	} else {
		database.SetterSchema(r.db, env.SchemaDBCl)
	}

	var reportData [][]string

	rows, err := r.db.Query(env.GetReportMassiveCl, yesterdayDate)

	if err != nil {
		log.Printf("cannot execute select query %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var reportClQuery domain.ReportOperationsMassiveCl

		err := rows.Scan(
			&reportClQuery.CustomerId,
			&reportClQuery.NumeroReg,
			&reportClQuery.ProccessName,
			&reportClQuery.FileName,
			&reportClQuery.Recipt,
			&reportClQuery.OperationNumber,
			&reportClQuery.LineDescription,
			&reportClQuery.Product,
			&reportClQuery.Subproduct,
			&reportClQuery.TemplateId,
			&reportClQuery.Categories,
			&reportClQuery.Flow,
			&reportClQuery.ShippingStatus,
			&reportClQuery.DatetimeSent,
			&reportClQuery.ReceivedStatus,
			&reportClQuery.OpeningDateTime,
			&reportClQuery.BounceType,
			&reportClQuery.BounceDateTime,
			&reportClQuery.BounceMotive,
		)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(reportClQuery)
		reportData = append(reportData, []string{
			reportClQuery.CustomerId.String,
			reportClQuery.NumeroReg.String,
			reportClQuery.ProccessName.String,
			reportClQuery.FileName.String,
			reportClQuery.Recipt.String,
			reportClQuery.OperationNumber.String,
			reportClQuery.LineDescription.String,
			reportClQuery.Product.String,
			reportClQuery.Subproduct.String,
			reportClQuery.TemplateId.String,
			reportClQuery.Categories.String,
			reportClQuery.Flow.String,
			reportClQuery.ShippingStatus.String,
			helper.DateToString(reportClQuery.DatetimeSent),
			helper.DateToString(reportClQuery.ReceivedStatus),
			helper.DateToString(reportClQuery.OpeningDateTime),
			reportClQuery.BounceType.String,
			helper.DateToString(reportClQuery.BounceDateTime),
			reportClQuery.BounceMotive.String,
		})
	}
	return reportData, nil
}
*/