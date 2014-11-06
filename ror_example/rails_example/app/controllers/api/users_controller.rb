module Api
  class UsersController < ApplicationController

    def index
        render json: User.all
    end
    def create
        print "Json is: "
        print params
        @u = User.new
        @u.username = params['username']
        if @u.save
            render json: @u
        else
            render nothing: true, status: :bad_request
        end
    end
    def show
        print params['id']
        @u = User.find_by_id(params['id'])
        render json: @u
    end
    def new
        create()
    end
    def destroy
        render json: {message: "coming soon"}
    end
    def edit
        @u = User.find_by_id(params['id'])
        @u.username = params['username']
        if @u.save
            render json: @u
        else
            render nothing: true, status: :bad_request
        end
    end
    def update
        edit()
    end
  end
end
