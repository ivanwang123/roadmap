import clsx from "clsx";
import React from "react";
import { useNotificationStore } from "../hooks/useNotificationStore";
import { NotificationType } from "../types";

type Props = {
  type?: "text" | "alert";
  showOnly?: NotificationType["type"];
  style?: string;
};

export function Notification({ type = "alert", showOnly, style }: Props) {
  const { notification, dismissNotification } = useNotificationStore();

  if (notification === null || (showOnly && showOnly !== notification?.type))
    return <></>;

  if (type === "text") {
    return <TextNotification notification={notification} style={style} />;
  } else {
    return (
      <AlertNotification
        notification={notification}
        dismissNotification={dismissNotification}
        style={style}
      />
    );
  }
}

type TextProps = {
  notification: NotificationType;
  style?: string;
};

function TextNotification({ notification, style }: TextProps) {
  return (
    <div
      className={clsx(
        style,
        notification?.type === "error" && "text-red-500",
        notification?.type === "success" && "text-green-500"
      )}
    >
      {notification?.message}
    </div>
  );
}

type AlertProps = TextProps & {
  dismissNotification: () => void;
};

function AlertNotification({
  notification,
  dismissNotification,
  style,
}: AlertProps) {
  return (
    <div
      className={clsx(
        "text-white px-5 py-2 rounded",
        style,
        notification?.type === "error" && "bg-red-400",
        notification?.type === "success" && "bg-green-400"
      )}
    >
      <button
        type="button"
        className="text-white mr-4"
        onClick={dismissNotification}
      >
        X
      </button>
      {notification?.message}
    </div>
  );
}
