import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import GroupIcon from '@mui/icons-material/Group';
import { ActionCard } from './ActionCard';

describe('ActionCard', () => {
    it('renders the label and calls onClick when clicked', async () => {
        const onClick = vi.fn();
        const user = userEvent.setup();
        render(<ActionCard icon={GroupIcon} label="עריכת נהגים" onClick={onClick} />);

        expect(screen.getByText('עריכת נהגים')).toBeInTheDocument();
        await user.click(screen.getByText('עריכת נהגים'));
        expect(onClick).toHaveBeenCalledTimes(1);
    });

    it('defaults to the non-primary variant', () => {
        render(<ActionCard icon={GroupIcon} label="ברירת מחדל" onClick={() => {}} />);
        expect(screen.getByText('ברירת מחדל')).toBeInTheDocument();
    });

    it('renders in the primary variant without crashing', async () => {
        const onClick = vi.fn();
        const user = userEvent.setup();
        render(<ActionCard icon={GroupIcon} label="פעולה ראשית" onClick={onClick} variant="primary" />);

        await user.click(screen.getByText('פעולה ראשית'));
        expect(onClick).toHaveBeenCalledTimes(1);
    });
});
