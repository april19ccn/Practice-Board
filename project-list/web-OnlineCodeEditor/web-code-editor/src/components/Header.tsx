import { ActionIcon } from "@mantine/core";
import { IconHeart } from "@tabler/icons-react";

const Header = () => {
    return (
        <>
            <div>
                <ActionIcon variant="gradient" size="xl" aria-label="Gradient action icon" gradient={{ from: "blue", to: "cyan", deg: 90 }}>
                    <IconHeart />
                </ActionIcon>
            </div>
        </>
    );
};

export default Header;