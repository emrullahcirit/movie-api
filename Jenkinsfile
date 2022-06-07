pipeline {
	agent any

	stages {
		stage('Docker Build') {
			steps {
				sh 'sudo docker build -t emrullahcirit/movie-api-image:latest .'
			}
		}
	}
	
}
