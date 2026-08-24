import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen, act } from '@testing-library/react';
import { GlobalStateProvider, useGlobalState } from './GlobalState';

function Probe() {
  const { level, setLevel } = useGlobalState();
  return (
    <button onClick={() => setLevel(3)}>{level}</button>
  );
}

describe('GlobalStateProvider', () => {
  beforeEach(() => {
    localStorage.clear();
  });

  it('defaults level to 0 when nothing is saved', () => {
    render(
      <GlobalStateProvider>
        <Probe />
      </GlobalStateProvider>
    );
    expect(screen.getByRole('button')).toHaveTextContent('0');
  });

  it('persists level changes to localStorage', async () => {
    render(
      <GlobalStateProvider>
        <Probe />
      </GlobalStateProvider>
    );
    await act(async () => {
      screen.getByRole('button').click();
    });
    expect(screen.getByRole('button')).toHaveTextContent('3');
    expect(localStorage.getItem('level')).toBe('3');
  });
});
