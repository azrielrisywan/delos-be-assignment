package config

import (
	supa "github.com/nedpals/supabase-go"
)

func SetupSupabase() *supa.Client {
	supabaseUrl := "https://lbwwemxfrzqvqmybxfdq.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imxid3dlbXhmcnpxdnFteWJ4ZmRxIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTUzMjE0OTEsImV4cCI6MjAzMDg5NzQ5MX0.sbIm2CyH30yx7V_e4JkFojCjMH6sf1qPCD_0b57TSOA"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	return supabase
}