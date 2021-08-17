import create from "zustand";
import { NotificationType } from "../types";

type NotificationStore = {
  notification: NotificationType | null;
  setNotification: (notification: NotificationType) => void;
  dismissNotification: () => void;
};

export const useNotificationStore = create<NotificationStore>((set) => ({
  notification: null,
  setNotification: (notification) => {
    set(() => ({ notification }));
  },
  dismissNotification: () =>
    set(() => ({
      notification: null,
    })),
}));
