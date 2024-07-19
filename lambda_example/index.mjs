import {
  S3Client,
  GetObjectCommand,
  PutObjectCommand,
} from "@aws-sdk/client-s3";
import sharp from "sharp";

const S3 = new S3Client();
const DESC_BUCKET = process.env.DESC_BUCKET;
const WIDTH = 200;
const FORMATS = {
  jpg: true,
  jpeg: true,
  png: true,
};

export const handler = async (event, context) => {
  const { eventTime, s3 } = event.Records[0];
  const src = s3.bucket.name;

  const key = decodeURIComponent(s3.object.key.replace(/\+/g, " "));
  const ext = key.replace(/^.*\./, "").toLowerCase();

  if (!FORMATS[ext]) {
    console.log("Invalid format");
    return;
  }

  try {
    const { Body, ContentType } = await S3.send(
      new GetObjectCommand({
        Bucket: src,
        Key: key,
      })
    );

    const image = await Body.transformToByteArray();
    const resizeImage = await sharp(image).resize(WIDTH).toBuffer();

    await S3.send(
      new PutObjectCommand({
        Bucket: DESC_BUCKET,
        Key: key,
        Body: resizeImage,
        ContentType,
      })
    );

    return {
      statusCode: 200,
      body: "Success",
      desc: DESC_BUCKET,
    };
  } catch (error) {
    console.log(error);
    return {
      statusCode: 400,
      body: "Failed",
      desc: DESC_BUCKET,
    };
  }
};
