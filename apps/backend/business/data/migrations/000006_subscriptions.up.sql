CREATE TABLE
  IF NOT EXISTS plans (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    lemonsqueezy_plan_id VARCHAR(255) UNIQUE NOT NULL,
    price NUMERIC(12, 2) NOT NULL CHECK (price > 0), -- Price stored with 2 decimal places, ensuring it's non-negative
    features JSONB, -- Store plan features as a JSON object (e.g., monitor limits, API access)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS subscriptions (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    team_id BIGINT NOT NULL REFERENCES teams (id),
    plan_id BIGINT NOT NULL REFERENCES plans (id),
    lemonsqueezy_subscription_id VARCHAR(255) UNIQUE NOT NULL, -- LemonSqueezy subscription ID
    plan_status VARCHAR(15) NOT NULL CHECK (
      plan_status IN (
        'active',
        'paused',
        'pending',
        'expired',
        'cancelled',
      )
    ), -- Subscription status
    current_period_end TIMESTAMP NOT NULL, -- Start of the current billing period
    current_period_start TIMESTAMP NOT NULL, -- End of the current billing period
    auto_renew BOOLEAN DEFAULT TRUE, -- Whether the subscription is auto-renewed,
    payment_method VARCHAR(50), -- Store the payment method used (e.g., card, PayPal),
    trial_period BOOLEAN DEFAULT FALSE, -- Indicate if the subscription is in a trial period
    trial_end TIMESTAMP, -- End of the trial period, if any
    total_spent NUMERIC(10, 2) DEFAULT 0, -- Total amount spent by the user/team on this subscription
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS billing_details (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    subscription_id BIGINT NOT NULL REFERENCES subscriptions (id) ON DELETE CASCADE,
    amount_paid NUMERIC(12, 2) NOT NULL CHECK (price > 0),
    currency VARCHAR(5) NOT NULL DEFAULT 'inr',
    payment_status VARCHAR(20) NOT NULL CHECK (
      payment_status IN ('paid', 'pending', 'failed', 'refunded')
    ),
    discount_amount NUMERIC(10, 2) DEFAULT 0, -- Discount applied to this payment
    tax_amount NUMERIC(10, 2) DEFAULT 0, -- Tax applied to this payment
    refund_amount NUMERIC(10, 2) DEFAULT 0, -- Amount refunded (if any)
    refund_date TIMESTAMP, -- Date of refund (if any)
    transaction_date TIMESTAMP NOT NULL,
  );

CREATE TABLE
  IF NOT EXISTS payment_history (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    billing_detail_id BIGINT NOT NULL REFERENCES billing_details (id) ON DELETE CASCADE,
    lemonsqueezy_invoice_id VARCHAR(255) UNIQUE NOT NULL, -- LemonSqueezy invoice ID
    lemonsqueezy_payment_id VARCHAR(255) UNIQUE NOT NULL, -- LemonSqueezy payment ID
    amount NUMERIC(10, 2) NOT NULL, -- Amount of the transaction
    payment_type VARCHAR(20) NOT NULL CHECK (
      payment_type IN ('charge', 'refund', 'adjustment')
    ), -- Type of payment (charge, refund, etc.)
    transaction_date TIMESTAMP NOT NULL, -- Date of the transaction
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );