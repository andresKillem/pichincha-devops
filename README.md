# Banco Pichincha Assessment DevOps : Node.js application to Amazon ECS

[![Run Status](https://api.shippable.com/projects/58f6fcddd1780a07007bba3f/badge?branch=master)](https://app.shippable.com/github/devops-recipes/deploy-ecs-basic) [![Coverage Badge](https://api.shippable.com/projects/58f6fcddd1780a07007bba3f/coverageBadge?branch=master)](https://app.shippable.com/github/devops-recipes/deploy-ecs-basic)

![AyeAye](https://github.com/devops-recipes/deploy-ecs-basic/blob/master/public/resources/images/captain.png)

This is  Node JS microservice with unit tests and coverage reports using mocha and istanbul. It also does a docker build once CI passes and then pushes the image to Amazon EC2 Container Registry

## Setup

To run this microservice architecture, you must first perform some setup steps to customize **shippable.yml** to point to your ECR and ECS accounts. 

### 1. Create an AWS integration

In order to allow Shippable to connect to your AWS services, you need to add an [integration](http://docs.shippable.com/platform/integration/overview/) through the Shippable UI. This integration will be referenced in the YAML as needed.

Log in to [Shippable](wwww.shippable.com) and follow instructions in the docs to [create an AWS keys integration](http://docs.shippable.com/platform/integration/aws-keys/). Note down the name of the integration you created.

### 2. Update YAML

Fork this repository to your account. 

All configuration is stored in the **shippable.yml** file at the root of this repository. You will need to make the following updates to customize the standard YAML included with this example:

* Change ECR_REPO (line 19) to point to your ECR repo. Please make sure the repo already exists on ECR.
* Update the `integrationName`(47) in the integration.hub section with the name of your AWS integration.
* Under `resources`, update the deploy-ecs-basic-image resource with the following:
   * `integration` should point to your AWS integration name
   * `sourceName` should point to the ECR repo
* Under `resources`, update the deploy-ecs-basic-ecs-cluster resource with the following:
   * `integration` should point to your AWS integration name
   * `sourceName` should point to your ECS cluster name. Please make sure the cluster already exists.
   * `region` should point to the region of your ECS cluster.

Commit your changes to source control.

## Run CI: Run some tests, build a Docker image and push to ECR

* [Enable your project for CI on Shippable](http://docs.shippable.com/ci/enable-project/) 
* Commit a small change to your repository (add a space or character in readme file). This will automatically trigger a build on Shippable. You can also trigger a manual build through the Shippable UI.
* The build will build your Docker image and push it to ECR.

## Enable the CD workflow: Deploy container to ECS

* Next, you need to enable config to deploy your Docker container to ECS. This logic is contained in the `jobs` and `resources` sections of your YAML, which is known as Assembly Line config in Shippable.
* [Add your Assembly Line](http://docs.shippable.com/platform/tutorial/workflow/add-assembly-line/). This will ensure that your `jobs` and `resources` sections are read.
* Rerun your CI job with another commit to the repository or with a manual build. The workflow should now trigger end to end and your container should be deployed to ECS.




