import { app, BrowserWindow, ipcMain } from "electron";
import path from "node:path";
import { makeClient } from "./grpc";

let win: BrowserWindow;
const client = makeClient();

async function createWindow() {
  win = new BrowserWindow({
    width: 960, height: 640,
    webPreferences: { preload: path.join(__dirname, "preload.js"), contextIsolation: true, nodeIntegration: false }
  });
  await win.loadURL("http://localhost:5173");
}

app.whenReady().then(createWindow);

ipcMain.handle("control:start", (_evt, profileId: string) => {
  return new Promise((resolve, reject) =>
    client.StartTunnel({ profile_id: profileId }, (err: unknown, resp: unknown) => err ? reject(err) : resolve(resp))
  );
});

ipcMain.handle("control:status", (_evt, profileId: string) => {
  return new Promise((resolve, reject) =>
    client.GetStatus({ id: profileId }, (err: unknown, resp: unknown) => err ? reject(err) : resolve(resp))
  );
});
