class CreateStories < ActiveRecord::Migration[5.2]
  def change
    create_table :stories, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_bigint, auto_increment: true
      t.text        :title, null: false
      t.text        :synopsis, null: false
      t.timestamps  null: false
    end
  end
end
