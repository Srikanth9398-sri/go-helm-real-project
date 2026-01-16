pipeline {
    agent any

    environment {
        IMAGE_NAME = "srikanth9398/go-app"
        DOCKER_CREDS = "dockerhub-creds"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/Srikanth9398-sri/go-helm-real-project.git'
            }
        }

        stage('Go Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t $IMAGE_NAME:${BUILD_NUMBER} .'
            }
        }

        stage('Docker Push') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: DOCKER_CREDS,
                    usernameVariable: 'DH_USER',
                    passwordVariable: 'DH_PASS'
                )]) {
                    sh '''
                      echo $DH_PASS | docker login -u $DH_USER --password-stdin
                      docker push $IMAGE_NAME:${BUILD_NUMBER}
                    '''
                }
            }
        }

        stage('Helm Deploy') {
            steps {
                sh '''
                  helm upgrade --install go-app ./helm-chart \
                    --set image.repository=$IMAGE_NAME \
                    --set image.tag=${BUILD_NUMBER}
                '''
            }
        }
    }

    post {
        success {
            echo "✅ Image pushed to Docker Hub and deployed"
        }
        failure {
            echo "❌ Pipeline failed"
        }
    }
}
