import { browser } from "$app/environment";

    // Increment tab count
    export function incrementTabCount(participantKey: string) {
      if (!browser) return;
      const currentVal = localStorage.getItem(participantKey);
      const count = parseInt(currentVal || '0', 10);
      console.log(`Incrementing tab count for ${participantKey}. Current before: ${currentVal}, Parsed: ${count}`);
      console.trace();
      localStorage.setItem(participantKey, String(count + 1));
      sessionStorage.setItem(`${participantKey}-tabIncremented`, "true");
    }

    // Decrement tab count
    export function decrementTabCount(participantKey: string) {
      if (!browser) return;
      const count = parseInt(localStorage.getItem(participantKey) || '0', 10);
      const newCount = Math.max(count - 1, 0);
      localStorage.setItem(participantKey, String(newCount));
    }