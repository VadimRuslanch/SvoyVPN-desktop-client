import { useEffect, useState } from "react";

declare global { interface Window { vadim: { start(id:string):Promise<any>; status(id:string):Promise<any>; } } }

export default function App() {
  const [state, setState] = useState<string>("DISCONNECTED");
  const pid = "demo-profile";

  useEffect(() => {
    const t = setInterval(async () => {
      const st = await window.vadim.status(pid);
      setState(st.state || "UNKNOWN");
    }, 1000);
    return () => clearInterval(t);
  }, []);

  return (
    <div style={{ padding: 24, fontFamily: "system-ui, sans-serif" }}>
      <h1>Вадим — Decktop</h1>
      <p>Состояние: <b>{state}</b></p>
      <button onClick={() => window.vadim.start(pid)}>Connect</button>
    </div>
  );
}
