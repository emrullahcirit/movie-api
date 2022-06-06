pipeline {
	agent any

	stages {
		stage('Docker Build') {
			steps {
				sh 'docker build -t emrullahcirit/movie-api-image:latest .'
			}
		}
	}
	
}
