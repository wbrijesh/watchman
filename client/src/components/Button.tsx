const PrimaryButton = ({
  text,
  onClick = () => {},
  disabled = false,
  type = "button",
  intent = "primary",
}: {
  text: string;
  onClick?: () => void;
  disabled?: boolean;
  type?: "button" | "submit" | "reset";
  intent?: "primary" | "danger";
}) => {
  return (
    <div>
      <button
        onClick={onClick}
        disabled={disabled}
        type={type}
        className={`text-white font-medium px-2 py-1 rounded ${
          intent === "primary"
            ? "bg-slate-700 disabled:bg-slate-500"
            : "bg-red-700 disabled:bg-red-500"
        }`}
      >
        {text}
      </button>
    </div>
  );
};

export { PrimaryButton };
