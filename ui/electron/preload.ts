import { contextBridge, ipcRenderer } from "electron";

contextBridge.exposeInMainWorld("vadim", {
  start: (profileId: string) => ipcRenderer.invoke("control:start", profileId),
  status: (profileId: string) => ipcRenderer.invoke("control:status", profileId)
});
