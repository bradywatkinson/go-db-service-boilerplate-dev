"""Update CommittedTimeStampMixin to GormModelMixin

Revision ID: f2ea8a1fc7b6
Revises: 2946eee19802
Create Date: 2018-10-29 10:10:39.349416

"""
from alembic import op
import sqlalchemy as sa
from sqlalchemy.dialects import postgresql

# revision identifiers, used by Alembic.
revision = 'f2ea8a1fc7b6'
down_revision = '2946eee19802'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('customer', sa.Column('created_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.add_column('customer', sa.Column('deleted_at', sa.DateTime(), nullable=True))
    op.add_column('customer', sa.Column('updated_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.create_index('ix_customer_deleted_at', 'customer', ['deleted_at'], unique=False)
    op.drop_column('customer', 'committed_timestamp')
    op.add_column('merchant', sa.Column('created_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.add_column('merchant', sa.Column('deleted_at', sa.DateTime(), nullable=True))
    op.add_column('merchant', sa.Column('updated_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.create_index('ix_merchant_deleted_at', 'merchant', ['deleted_at'], unique=False)
    op.drop_column('merchant', 'committed_timestamp')
    op.add_column('product', sa.Column('created_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.add_column('product', sa.Column('deleted_at', sa.DateTime(), nullable=True))
    op.add_column('product', sa.Column('updated_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.create_index('ix_product_deleted_at', 'product', ['deleted_at'], unique=False)
    op.drop_column('product', 'committed_timestamp')
    op.add_column('review', sa.Column('created_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.add_column('review', sa.Column('deleted_at', sa.DateTime(), nullable=True))
    op.add_column('review', sa.Column('updated_at', sa.DateTime(), server_default=sa.text('now()'), nullable=False))
    op.create_index('ix_review_deleted_at', 'review', ['deleted_at'], unique=False)
    op.drop_column('review', 'committed_timestamp')
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('review', sa.Column('committed_timestamp', postgresql.TIMESTAMP(), server_default=sa.text('now()'), autoincrement=False, nullable=False))
    op.drop_index('ix_review_deleted_at', table_name='review')
    op.drop_column('review', 'updated_at')
    op.drop_column('review', 'deleted_at')
    op.drop_column('review', 'created_at')
    op.add_column('product', sa.Column('committed_timestamp', postgresql.TIMESTAMP(), server_default=sa.text('now()'), autoincrement=False, nullable=False))
    op.drop_index('ix_product_deleted_at', table_name='product')
    op.drop_column('product', 'updated_at')
    op.drop_column('product', 'deleted_at')
    op.drop_column('product', 'created_at')
    op.add_column('merchant', sa.Column('committed_timestamp', postgresql.TIMESTAMP(), server_default=sa.text('now()'), autoincrement=False, nullable=False))
    op.drop_index('ix_merchant_deleted_at', table_name='merchant')
    op.drop_column('merchant', 'updated_at')
    op.drop_column('merchant', 'deleted_at')
    op.drop_column('merchant', 'created_at')
    op.add_column('customer', sa.Column('committed_timestamp', postgresql.TIMESTAMP(), server_default=sa.text('now()'), autoincrement=False, nullable=False))
    op.drop_index('ix_customer_deleted_at', table_name='customer')
    op.drop_column('customer', 'updated_at')
    op.drop_column('customer', 'deleted_at')
    op.drop_column('customer', 'created_at')
    # ### end Alembic commands ###
