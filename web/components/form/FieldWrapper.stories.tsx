import { Meta, Story } from "@storybook/react";
import { FieldWrapper } from "./FieldWrapper";

export default {
  title: "Components/FieldWrapper",
  component: FieldWrapper,
} as Meta;

const Template: Story = (args: any) => <FieldWrapper {...args} />;

export const Input = Template.bind({});
Input.args = {
  id: "id",
  label: "Label",
  error: undefined,
  children: <input type="text" />,
};
