<template>
    <div class="item">
        <i>
            <slot name="icon"></slot>
        </i>
        <div class="details">
            <h3>
                <slot name="heading"></slot>
            </h3>
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>
.item {
    padding-top: 2rem;
    display: flex;
    position: relative;
}

.details {
    flex: 1;
    margin-left: 1rem;
}

i {
    display: flex;
    place-items: center;
    place-content: center;
    color: var(--color-text);
    border: 1px solid var(--color-border);
    background: var(--color-background);
    border-radius: 8px;
    width: 50px;
    height: 50px;
}

h3 {
    font-size: 1.2rem;
    font-weight: 500;
    margin-bottom: 0.4rem;
    color: var(--color-heading);
}
</style>
