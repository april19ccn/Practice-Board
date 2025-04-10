import { create } from "zustand";
import { produce } from "immer";

type State = {
    firstName: string;
    lastName: string;
    deep: {
        nested: {
            obj: { count: number };
        };
    };
};

type Action = {
    updateFirstName: (firstName: State["firstName"]) => void;
    updateLastName: (lastName: State["lastName"]) => void;
    testInc: () => void;
    immerInc: () => void;
    normalInc: () => void;
};

// Create your store, which includes both state and (optionally) actions
const usePersonStore = create<State & Action>((set) => ({  // set 是在 node_modules/zustand/esm/vanilla.mjs 中 const initialState = state = createState(setState, getState, api);
    firstName: "",
    updateFirstName: (firstName) => set(() => ({ firstName: firstName })),

    lastName: "",
    updateLastName: (lastName) => set(() => ({ lastName: lastName })),

    deep: { nested: { obj: { count: 1 } } },
    testInc: () => set(() => ({ deep: { nested: { obj: { count: 2 } } } })),
    immerInc: () =>
        set(
            produce((state: State) => {
                ++state.deep.nested.obj.count;
            })
        ),
    // immerInc: () =>
    //     set((state) =>
    //         produce(state, (state: State) => {
    //             ++state.deep.nested.obj.count;
    //         })
    //     ),
    normalInc: () =>
        set((state) => ({
            deep: {
                ...state.deep,
                nested: {
                    ...state.deep.nested,
                    obj: {
                        ...state.deep.nested.obj,
                        count: state.deep.nested.obj.count + 1,
                    },
                },
            },
        })),
}));

// In consuming app
function TestUpdating() {
    // "select" the needed state and actions, in this case, the firstName value
    // and the action updateFirstName
    const firstName = usePersonStore((state) => state.firstName);
    const updateFirstName = usePersonStore((state) => state.updateFirstName);

    const lastName = usePersonStore((state) => state.lastName);
    const updateLastName = usePersonStore((state) => state.updateLastName);

    const deep = usePersonStore((state) => state.deep);
    const testInc = usePersonStore((state) => state.testInc);
    const immerInc = usePersonStore((state) => state.immerInc);
    const normalInc = usePersonStore((state) => state.normalInc);

    return (
        <main>
            <label>
                First name:
                <input
                    // Update the "firstName" state
                    onChange={(e) => updateFirstName(e.currentTarget.value)}
                    value={firstName}
                />
                Last Name:
                <input
                    // Update the "lastName" state
                    onChange={(e) => updateLastName(e.currentTarget.value)}
                    value={lastName}
                />
                <div>
                    <button onClick={testInc} style={{ marginRight: "1rem" }}>
                        Test Inc
                    </button>
                    <button onClick={immerInc} style={{ marginRight: "1rem" }}>
                        Immer Inc
                    </button>
                    <button onClick={normalInc} style={{ marginRight: "1rem" }}>
                        Normal Inc
                    </button>
                </div>
            </label>

            <p>
                Hello,{" "}
                <strong>
                    {firstName}!{lastName}
                </strong>
                <div>{JSON.stringify(deep)}</div>
            </p>
        </main>
    );
}

export default TestUpdating;


