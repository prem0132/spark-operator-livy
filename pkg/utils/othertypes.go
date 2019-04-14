package utils

// SparkAppSpec is the json Body to be sent to the spark operator
type SparkAppSpec struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Type                string   `json:"type"`
		Mode                string   `json:"mode"`
		Image               string   `json:"image"`
		ImagePullPolicy     string   `json:"imagePullPolicy"`
		MainClass           string   `json:"mainClass"`
		MainApplicationFile string   `json:"mainApplicationFile"`
		Arguments           []string `json:"arguments"`
		RestartPolicy       struct {
			Type string `json:"type"`
		} `json:"restartPolicy"`
		HadoopConf struct {
			FsS3AAccessKey string `json:"fs.s3a.access.key"`
			FsS3ASecretKey string `json:"fs.s3a.secret.key"`
			FsS3AImpl      string `json:"fs.s3a.impl"`
		} `json:"hadoopConf"`
		Driver struct {
			Labels struct {
				Version string `json:"version"`
			} `json:"labels"`
			ServiceAccount string `json:"serviceAccount"`
			EnvVars        struct {
				CHECKPOINTDIR string `json:"CHECKPOINT_DIR"`
			} `json:"envVars"`
			EnvSecretKeyRefs struct {
				AWSACCESSKEYID struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"AWS_ACCESS_KEY_ID"`
				AWSSECRETACCESSKEY struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"AWS_SECRET_ACCESS_KEY"`
			} `json:"envSecretKeyRefs"`
		} `json:"driver"`
		Executor struct {
			Instances int `json:"instances"`
			Labels    struct {
				Version string `json:"version"`
			} `json:"labels"`
			EnvSecretKeyRefs struct {
				AWSACCESSKEYID struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"AWS_ACCESS_KEY_ID"`
				AWSSECRETACCESSKEY struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"AWS_SECRET_ACCESS_KEY"`
			} `json:"envSecretKeyRefs"`
		} `json:"executor"`
	} `json:"spec"`
}
