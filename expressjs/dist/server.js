"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
const app_1 = __importDefault(require("./app"));
function bootServer(port) {
    return app_1.default.listen(PORT, () => {
        console.log(`API server listening on port ${port}`);
    });
}
const PORT = parseInt((_a = process.env["PORT"]) !== null && _a !== void 0 ? _a : "5005", 10);
void bootServer(PORT);
