interface ButtonProps {
    disabled?: boolean;
    title: string;
    onClick: () => void;
}

const Button = ({ title, onClick }: ButtonProps) => {
    return (
        <button style={{ maxWidth: "140px", minWidth: "80px", height: "30px" }} onClick={onClick}>
            {title}
        </button>
    );
};

export default Button;
