import express, {
  urlencoded,
  json,
  Application,
  Request,
  Response,
} from "express";

function buildApp(): Application {
  const app: Application = express();

  app.use(urlencoded({ extended: true }));
  app.use(json());
  app.set("trust proxy", 1);
  app.use("/ping", (_: Request, res: Response) => {
    console.log(process.env);
    res.status(200).json({ message: "pong" });
  });

  return app;
}

export default buildApp();
