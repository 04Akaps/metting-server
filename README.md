<h1> Metting </h1>

어플리케이션을 가정한 개인 사이드 프로젝트

1. API

일반적인 서버 API

2. Internal

일부 수치 및 미처리 파일을 관리하기 위한 cron 모듈


<h1> S3 </h1>

유저의 이미지 파일을 관리하기 위해 사용

<h3> 정책 관리 </h3>

```
{
    "Version": "2012-10-17",
    "Id": "Policy1717656453611",
    "Statement": [
        {
            "Sid": "Stmt1717656451362",
            "Effect": "Allow",
            "Principal": {
                "AWS": "적당한 IAM"
            },
            "Action": "s3:*",
            "Resource": "적당한 ARN"
        },
        {
            "Sid": "PublicRead",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "적당한 버킷 경로"
        }
    ]
}    
```

상황에 따라서 해당 정책은 수정이 가능하다.
이미지 캐시를 위하여 프로젝트에서는 CloudFront를 적용 하였다.

<h3> CloudFront 정책 </h3>

```
{
    "Version": "2008-10-17",
    "Id": "PolicyForCloudFrontPrivateContent",
    "Statement": [
        {
            "Sid": "AllowCloudFrontServicePrincipal",
            "Effect": "Allow",
            "Principal": {
                "Service": "cloudfront.amazonaws.com"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::metting-s3/*",
            "Condition": {
                "StringEquals": {
                  "AWS:SourceArn": "적당한 ARN"
                }
            }
        }
    ]
  }
```


<h1> MySQL Schema </h1>

```
CREATE table IF NOT EXISTS `user` (
    `t_id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_name` VARCHAR(100) UNIQUE NOT NULL,
    `image` JSON DEFAULT("[]") NOT NULL,
    `description` VARCHAR(300) DEFAULT "",
    `hobby` JSON DEFAULT ("[]") NOT NULL,
    `is_valid` TINYINT(1) DEFAULT false
);


CREATE table IF NOT EXISTS `user_location` (
    `user_name` VARCHAR(100) UNIQUE NOT NULL,
    `latitude`  DOUBLE NULL COMMENT '위도',
    `hardness` DOUBLE NOT NULL COMMENT '경도',
    `location` POINT NOT NULL COMMENT '위치',
    
    SPATIAL INDEX `spatial_index` (`location`)
);


CREATE TABLE IF NOT EXISTS `user_like` (
    `from_user` VARCHAR(100) NOT NULL,
    `to_user` VARCHAR(100) NOT NULL,
    `status` ENUM("send", "checked", "fail") NOT NULL,
    `created_time` BIGINT NOT NULL,
    `updated_time` BIGINT DEFAULT 0 
);
```