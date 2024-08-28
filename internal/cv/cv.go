package cv

type CV struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
	JobTitles []struct {
		Title string   `yaml:"title"`
		Tags  []string `yaml:"tags"`
	} `yaml:"job_titles"`
}
