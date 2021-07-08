import React from "react";

type Props = {
  icon: any;
  size: number;
};

function Icon({ icon: Icon, size }: Props) {
  return (
    <span className="bg-secondary p-2 rounded-sm">
      <Icon className="fill-current" width={size} height={size} />
    </span>
  );
}

export default Icon;
