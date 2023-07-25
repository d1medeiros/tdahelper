

````mermaid
sequenceDiagram
    user->>wallet: search for events
    wallet->>eventmanager: select all for mounth x
    eventmanager-->>wallet: return
    wallet-->>user: return
````