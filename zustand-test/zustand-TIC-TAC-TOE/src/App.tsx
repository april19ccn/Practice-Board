// TEST-COMPONENTS
import Board from "./components/Board";

// TEST-EXAMPLES
import TestUpdating from "./examples/TestUpdating";

// TEST-THIRD
import TestImmer from "./test-ex/immer";

function App() {
    console.log("TestImmer------", TestImmer);

    return (
        <>
            <Board />

            <TestUpdating></TestUpdating>
        </>
    );
}

export default App;
