pipeline {
    agent any

    environment {
        IMAGE_NAME = "auth-service"
        IMAGE_TAG  = "latest"
    }

    options {
        timestamps()
        timeout(time: 30, unit: 'MINUTES')
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/Oumaymazekri/projetannuel.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'npm install'
            }
        }

        stage('Lint') {
            steps {
                sh 'npx eslint . || true'
            }
        }

        stage('Tests') {
            steps {
                sh 'npm test || true'
            }
        }

        stage('SonarQube Analysis') {
            environment {
                SONAR_TOKEN = credentials('sonar-token')
            }
            steps {
                sh """
                npx sonar-scanner \
                -Dsonar.projectKey=auth-service \
                -Dsonar.sources=. \
                -Dsonar.host.url=http://localhost:9000 \
                -Dsonar.login=$SONAR_TOKEN
                """
            }
        }

        stage('Docker Build') {
            steps {
                sh "docker build -t $IMAGE_NAME:$IMAGE_TAG ."
            }
        }

        stage('Deploy') {
            environment {
                JWT_SECRET = credentials('jwt-secret')
                EMAIL_PASS = credentials('email-password')
                EMAIL_USER = credentials('email-user')
            }
            steps {
                sh """
                docker stop auth-service || true
                docker rm auth-service || true

                docker run -d \
                --name auth-service \
                -p 3000:3000 \
                -e JWT_SECRET_KEY=$JWT_SECRET \
                -e EMAIL_USER=$EMAIL_USER \
                -e EMAIL_PASSWORD=$EMAIL_PASS \
                $IMAGE_NAME:$IMAGE_TAG
                """
            }
        }
    }

    post {
        success {
            withCredentials([string(credentialsId: 'slack-webhook', variable: 'SLACK_URL')]) {
                sh """
                curl -X POST -H 'Content-type: application/json' \
                --data '{"text":"✅ Pipeline AUTH-SERVICE SUCCESS"}' \
                $SLACK_URL
                """
            }
        }
        failure {
            withCredentials([string(credentialsId: 'slack-webhook', variable: 'SLACK_URL')]) {
                sh """
                curl -X POST -H 'Content-type: application/json' \
                --data '{"text":"❌ Pipeline AUTH-SERVICE FAILED"}' \
                $SLACK_URL
                """
            }
        }
    }
}
