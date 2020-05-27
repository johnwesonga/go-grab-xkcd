Go-Grab-XKCD
=============
Golang http app that grabs random XKCD strip using the XKCD api.

Original code: https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html

Built this to test GCP Cloud Run

Deployment Steps
=================
Assumes you have gcloud CLI installed

- gcloud builds submit --tag gcr.io/{{PROJECT-ID}}/go-grab-xkcd
- gcloud run deploy --image gcr.io/{{PROJECT-ID}}/go-grab-xkcd --platform managed
