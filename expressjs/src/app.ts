import express, {urlencoded, json, Application} from "express";

function buildApp(): Application {
  const app: Application = express();

  app.use(urlencoded({ extended: true }));
  app.use(json());
  app.set("trust proxy", 1);

  return app;
}

export default buildApp();
