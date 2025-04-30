import { MantineProvider } from "@mantine/core";
import { theme } from "./theme";

import Core from "./views/Core"

function App() {

    return (
        <MantineProvider theme={theme} defaultColorScheme="light">
            <Core></Core>
        </MantineProvider>
    );
}

export default App;
